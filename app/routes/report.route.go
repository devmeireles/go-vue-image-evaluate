package routes

import (
	"github.com/devmeireles/go-vue-image-evaluate/app/controllers"
	_ "github.com/devmeireles/go-vue-image-evaluate/app/docs"
	"github.com/gofiber/fiber/v2"
)

func ReportRoutes(app *fiber.App) {
	report := app.Group("/report")
	report.Post("/", controllers.CreateReport)
}
