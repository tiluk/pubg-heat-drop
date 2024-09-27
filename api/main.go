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
	lobbyService, sessionService := initServices()

	registerRoutes(app, lobbyService, sessionService)

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

func initServices() (*lobby.Service, *session.Service) {
	lobbyRepository := lobby.NewRepository(cache)
	sessionRepository := session.NewRepository(cache)

	lobbyService := lobby.NewService(lobbyRepository)
	sessionService := session.NewService(sessionRepository)

	return lobbyService, sessionService
}

func registerRoutes(app *fiber.App, lobbyService *lobby.Service, sessionService *session.Service) {
	lobbyController := lobby.NewController(lobbyService)
	sessionController := session.NewController(sessionService)

	routes := app.Group("/api")

	routes.Post("/lobby", lobbyController.PostLobby)
	routes.Get("/lobby/:id", lobbyController.GetLobby)
	routes.Post("/session", sessionController.PostSession)
	routes.Get("/session/:id", sessionController.GetSession)
}
