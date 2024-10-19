package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tiluk/pubg-heat-drop/auth"
	"github.com/tiluk/pubg-heat-drop/lobby"
	"github.com/tiluk/pubg-heat-drop/session"
)

func main() {
	app := fiber.New()

	initEnv()
	cache := initCache()
	sessionService := initSession(cache)
	lobbyService := initServices(cache, sessionService)

	setupMiddlewares(app, sessionService)

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
	cache := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       0,
	})

	return cache
}

func initSession(cache *redis.Client) *session.SessionService {
	sessionRepository := session.NewRepository(cache)

	return session.NewService(sessionRepository)
}

func initServices(cache *redis.Client, sessionService *session.SessionService) *lobby.LobbyService {
	lobbyRepository := lobby.NewRepository(cache)

	lobbyService := lobby.NewService(lobbyRepository, sessionService)

	return lobbyService
}

func setupMiddlewares(app *fiber.App, sessionService *session.SessionService) {
	app.Use(auth.NewAuthMiddleware(sessionService))
}

func registerRoutes(app *fiber.App, lobbyService *lobby.LobbyService, sessionService *session.SessionService) {
	lobbyController := lobby.NewController(lobbyService)
	sessionController := session.NewController(sessionService)

	routes := app.Group("/api")

	routes.Post("/auth", sessionController.PostSession)
	routes.Post("/lobby", lobbyController.PostLobby)
	routes.Get("/lobby/:id", lobbyController.GetLobby)
	routes.Post("/lobby/:id/vote", lobbyController.PostLobbyVote)
}
