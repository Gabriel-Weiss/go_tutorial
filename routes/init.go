package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", HandleCardList)
	app.Get("/shop", HandleShop)
	app.Get("/contact", HandleContact)
	app.Get("/about", HandleAbout)
	app.Post("/search", HandleShop)
}
