package handlers

import (
	"fmt"

	"github.com/afnank19/theaegean-go/services"
	"github.com/gofiber/fiber/v2"
)

func GetABlog(c *fiber.Ctx) error {

	blogId := c.Params("id")
	fmt.Println(blogId)

	response := services.FetchABlogById(blogId)
	
	return c.JSON(response)
}

func GetAllBlogs(c *fiber.Ctx) error {
	return c.SendString("UNIMPLEMENTED: Fetch all blogs (paginated)")
}

func PostBlog(c *fiber.Ctx) error {
	return c.SendString("UNIMPLEMENTED: Add blog data to db")
}

func GetComments(c *fiber.Ctx) error {

	blogId := c.Params("id")

	lastDocId := c.Query("lastDocId")
	fmt.Println("handler: " + lastDocId)

	response, err := services.FetchBlogComments(blogId, lastDocId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "I did this on purpose")
	}


	fmt.Println(response)

	return c.JSON(response)
}

func PostComment(c *fiber.Ctx) error {
	return c.SendString("UNIMPLEMENTED: Add comment data to db")
}