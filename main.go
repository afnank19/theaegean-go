package main

import (
	"github.com/afnank19/theaegean-go/config"

	"github.com/afnank19/theaegean-go/routes"
	"github.com/gofiber/fiber/v2"
)

// New beginnings
func main() {
	PORT := ":8081"

	config.InitializeFirebase("/home/afnank19/Desktop/development/personal/go-theaegean/secrets/the-sapphire-19ba8-firebase-adminsdk-fu3od-d725600a17.json")

	// Fiber setup
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("We're up alive and healthy")
	})

	api := app.Group("/api")
	api = routes.GroupBlogRoute(api)

	app.Listen(PORT)
}



