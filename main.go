package main

import (
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitViperEnvConfig()
	configs.InitCache()
	initFiberApp()
}

func initFiberApp() {
	app := fiber.New()
	apiRouter := app.Group("/api")

	routes.InitApiRoute(apiRouter)

	serverPort := configs.ViperEnv.Get("SERVER_PORT").(string)
	err := app.Listen(":" + serverPort)

	if err != nil {
		panic(err)
	}
}
