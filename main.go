package main

import (
	"log"

	"github.com/Gabriel-Weiss/go_tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./assets")
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
