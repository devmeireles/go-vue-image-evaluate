package main

import (
	database "github.com/devmeireles/go-vue-image-evaluate/app/config"
	"github.com/devmeireles/go-vue-image-evaluate/app/routes"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes() {
	app := fiber.New()
	routes.ReportRoutes(app)
	app.Listen(":3000")
}

func setupDatabase() {
	database.ConnectDb()
}
