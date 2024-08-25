package main

import (
	"log"

	"github.com/Gabriel-Weiss/go_tutorial/views/pages"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {

	app := fiber.New()
	app.Get("/", adaptor.HTTPHandler(templ.Handler(pages.IndexTempl("Hello world!!!"))))

	log.Fatal(app.Listen(":3000"))
}
