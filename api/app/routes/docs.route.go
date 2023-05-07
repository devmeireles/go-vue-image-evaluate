package routes

import (
	_ "github.com/devmeireles/go-vue-image-evaluate/app/docs"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func DocsRoutes(app *fiber.App) {
	app.Get("/docs/*", fiberSwagger.WrapHandler)
}
