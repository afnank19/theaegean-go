package routes

import (
	"fmt"

	"github.com/afnank19/theaegean-go/handlers"
	"github.com/gofiber/fiber/v2"
)

// Althought this file is called blog routes
// I may just add all routes in this file, but that
// is to be decided

func GroupBlogRoute(api fiber.Router) fiber.Router {
	blogRoutes := api.Group("/blogs", temporaryLogger)

	blogRoutes.Get("/:id", handlers.GetABlog)
	blogRoutes.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("UNIMPLEMENTED: Get all blogs")
	})

	return blogRoutes
}

// Remove when necessary
func temporaryLogger(c *fiber.Ctx)  error {
	fmt.Println("Im just a middleware")
	return c.Next()
}