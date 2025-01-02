package services

import (
	"context"
	"log"

	"github.com/afnank19/theaegean-go/config"
)

func FetchABlogById(blogContentId string) map[string]interface{}  {
	db, err := config.GetFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}

	// Handle this error !!!!
	dSnap, _ := db.Collection("blogContent").Doc(blogContentId).Get(context.Background())

	blogContent := dSnap.Data()

	return blogContent;


}