package routes

import (
	"github.com/gofiber/fiber/v2"
)

// todo: Make use of FakeShopAPI in requests
func SetupRoutes(app *fiber.App) {
	app.Get("/", HandleCardList)
	app.Get("/shop", HandleShop)
	app.Get("/shop/:id<int>", HandleDetails)
	app.Get("/contact", HandleContact)
	app.Get("/cart", HandleCart)
	app.Get("/about", HandleAbout)
	app.Post("/search", HandleShop)
	app.Post("/category", HandleCategories)
}
