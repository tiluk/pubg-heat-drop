package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tiluk/pubg-heat-drop/lobby"
	"github.com/tiluk/pubg-heat-drop/session"
)

var cache *redis.Client

func main() {
	app := fiber.New()

	initEnv()
	initCache()
	lobbyHandler, sessionHandler := initHandlers()

	registerRoutes(app, lobbyHandler, sessionHandler)

	log.Fatal(app.Listen(":8080"))
}

func initEnv() {
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func initCache() *redis.Client {
	cache = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       0,
	})

	return cache
}

func initHandlers() (*lobby.LobbyHandler, *session.SessionHandler) {
	lobbyHandler := lobby.NewLobbyHandler(cache)
	sessionHandler := session.NewSessionHandler(cache)

	return lobbyHandler, sessionHandler
}

func registerRoutes(app *fiber.App, lobbyHandler *lobby.LobbyHandler, sessionHandler *session.SessionHandler) {
	routes := app.Group("/api")

	routes.Post("/lobby", lobbyHandler.CreateLobby)
	routes.Get("/lobby/:id", lobbyHandler.GetLobby)
	routes.Post("/session", sessionHandler.CreateSession)
	routes.Get("/session/:id", sessionHandler.GetSession)
}
