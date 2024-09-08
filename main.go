package main

import (
	"log"

	"github.com/Gabriel-Weiss/go_tutorial/views/pages"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./assets")
	app.Get("/", adaptor.HTTPHandler(templ.Handler(pages.IndexPage("Hello world!!!"))))
	app.Get("/shop", adaptor.HTTPHandler(templ.Handler(pages.ShopPage("Hello world!!!"))))
	app.Get("/categories", adaptor.HTTPHandler(templ.Handler(pages.CategoriesPage("Hello world!!!"))))
	app.Get("/contact", adaptor.HTTPHandler(templ.Handler(pages.ContactPage("Hello world!!!"))))
	app.Get("/about", adaptor.HTTPHandler(templ.Handler(pages.AboutPage("Hello world!!!"))))

	log.Fatal(app.Listen(":3000"))
}
