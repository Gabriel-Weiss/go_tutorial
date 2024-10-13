package routes

import (
	"github.com/Gabriel-Weiss/go_tutorial/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const BaseUrl = "https://fakestoreapi.com"

var store = session.New(session.Config{
	CookieSessionOnly: true,
})

// TODO: Implement the commented functions
func SetupRoutes(app *fiber.App) {
	app.Get("/products", HandleAllProducts)
	app.Get("/products/category/:category", HandleProductsInCategory)
	app.Get("/products/:id", HandleProduct)
	app.Get("/categories", HandleAllCategories)
	app.Get("/signup", HandleSignUp)
	app.Post("/signup", HandlePostSignUp)
	app.Get("/signin", HandleSignIn)
	app.Post("/signin", HandlePostSignIn)
	app.Get("/contact", HandleContact)
	app.Get("/about", HandleAbout)
	// app.Post("/search", HandleShop)
	app.Get("/signout", HandleSignOut)

	protected := app.Group("/shopee", AuthenticationMiddleware)
	protected.Get("/users", HandleAllUsers)
	protected.Get("/users/:id", HandleUser)
	protected.Get("/carts", HandleAllCarts)
	// protected.Get("/carts/:id", Handle...)
	// protected.Get("/carts/user/:id", Handle...)

	app.Get("*", HandleNotFound)
}

func AuthenticationMiddleware(c *fiber.Ctx) error {
	session, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	}

	token, ok := session.Get("token").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	userID, ok := session.Get("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	decryptedUserID, err := services.Decrypt(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to decrypt user ID")
	}

	c.Locals("user_id", decryptedUserID)
	c.Locals("token", token)

	return c.Next()

}
