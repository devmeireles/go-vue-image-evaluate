package main

import (
	database "github.com/devmeireles/go-vue-image-evaluate/app/config"
	"github.com/devmeireles/go-vue-image-evaluate/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.ReportRoutes(app)
	routes.DocsRoutes(app)
	app.Listen(":3000")
}

func setupDatabase() {
	database.ConnectDb()
}
