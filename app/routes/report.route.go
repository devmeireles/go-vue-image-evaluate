package routes

import (
	"github.com/devmeireles/go-vue-image-evaluate/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ReportRoutes(app *fiber.App) {
	report := app.Group("/report")
	report.Post("/", controllers.CreateReport)
}
