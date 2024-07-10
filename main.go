package main

import (
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/handlers"
	"github.com/Strayneko/KomikcastAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitViperEnvConfig()
	configs.InitLogger()
	configs.InitCache()
	initFiberApp()
}

func initFiberApp() {
	handler := handlers.NewHandler()
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})
	apiRouter := app.Group("/api")

	routes.InitApiRoute(apiRouter)

	serverPort := configs.ViperEnv.Get("SERVER_PORT").(string)
	err := app.Listen(":" + serverPort)

	if err != nil {
		panic(err)
	}
}
