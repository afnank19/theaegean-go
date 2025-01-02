package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/afnank19/theaegean-go/config"
	"github.com/gofiber/fiber/v2"
)

func GetABlog(c *fiber.Ctx) error {

	blogId := c.Params("id")
	fmt.Println(blogId)

	fStoreClient, err := config.GetFirestoreClient()
	if err != nil {
		 log.Fatalln(err)
	}

	dSnap, _ := fStoreClient.Collection("blogContent").Doc(blogId).Get(context.Background())

	m := dSnap.Data()
	return c.JSON(m)
}