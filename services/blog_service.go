package services

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
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

func FetchBlogComments(blogContentId, lastDocId string) ([]*firestore.DocumentSnapshot, error){
	db, err := config.GetFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}

	if lastDocId == "" {
		// Split this up, horrible readability
		docSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").OrderBy("postDate", firestore.Desc).Limit(5).Documents(context.Background()).GetAll()

		for _, doc := range docSnap {
			fmt.Println(doc.Data())
		}

		return docSnap, nil
	}

	lastDocSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").Doc(lastDocId).Get(context.Background())

	// Split this up, horrible readability
	docSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").OrderBy("postDate", firestore.Desc).Limit(5).StartAfter(lastDocSnap).Documents(context.Background()).GetAll()

	// Add this to a map or something to send back, not just print it
	for _, doc := range docSnap {
		fmt.Println(doc.Data())
	}

	return docSnap, nil
}