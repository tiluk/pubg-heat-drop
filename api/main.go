package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tiluk/pubg-heat-drop/lobby"
	"github.com/tiluk/pubg-heat-drop/session"
	"github.com/tiluk/pubg-heat-drop/vote"
)

var cache *redis.Client

func main() {
	app := fiber.New()

	initEnv()
	initCache()
	lobbyService, sessionService, voteService := initServices()

	registerRoutes(app, lobbyService, sessionService, voteService)

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

func initServices() (*lobby.LobbyService, *session.SessionService, *vote.VoteService) {
	lobbyRepository := lobby.NewRepository(cache)
	sessionRepository := session.NewRepository(cache)

	lobbyService := lobby.NewService(lobbyRepository)
	sessionService := session.NewService(sessionRepository)

	voteService := vote.NewService(sessionService, lobbyService)

	return lobbyService, sessionService, voteService
}

func registerRoutes(app *fiber.App, lobbyService *lobby.LobbyService, sessionService *session.SessionService, voteService *vote.VoteService) {
	lobbyController := lobby.NewController(lobbyService)
	sessionController := session.NewController(sessionService)
	voteController := vote.NewController(voteService)

	routes := app.Group("/api")

	routes.Post("/lobby", lobbyController.PostLobby)
	routes.Get("/lobby/:id", lobbyController.GetLobby)
	routes.Post("/session", sessionController.PostSession)
	routes.Get("/session/:id", sessionController.GetSession)
	routes.Post("/vote/:id", voteController.PostVote)
}
