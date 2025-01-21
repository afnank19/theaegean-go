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

func FetchBlogComments(blogContentId, lastDocId string) ([]map[string]interface{}, error){
	db, err := config.GetFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}

	if lastDocId == "" {
		// Split this up, horrible readability
		docSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").OrderBy("postDate", firestore.Desc).Limit(5).Documents(context.Background()).GetAll()

		// can be a helper function
		var results []map[string]interface{}
		for _, doc := range docSnap {
			data := make(map[string]interface{})
			if err := doc.DataTo(&data); err != nil {
				log.Printf("failed to unmarshal document %s: %v", doc.Ref.ID, err)
				continue
			}
			results = append(results, data)
		}
	
		return results, nil
	}

	lastDocSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").Doc(lastDocId).Get(context.Background())

	// Split this up, horrible readability
	docSnap, _ := db.Collection("blogContent").Doc(blogContentId).Collection("comments").OrderBy("postDate", firestore.Desc).Limit(5).StartAfter(lastDocSnap).Documents(context.Background()).GetAll()

	// Add this to a map or something to send back, not just print it
	var results []map[string]interface{}
	for _, doc := range docSnap {
		data := make(map[string]interface{})
		if err := doc.DataTo(&data); err != nil {
			log.Printf("failed to unmarshal document %s: %v", doc.Ref.ID, err)
			continue
		}
		results = append(results, data)
	}

	return results, nil
}

func FetchUserBlogs(userId, lastDocId string) ([]map[string]interface{}, error){
	fmt.Println("UNIMPLEMENTED: Fetch blogs by a user(Paginated)")

	db, err := config.GetFirestoreClient()
	if err != nil {
		log.Fatalln(err)
	}

	if lastDocId == "" {
		docSnap, _ := db.Collection("blogMeta").Where("authorId", "==", userId).OrderBy("postDate", firestore.Desc).Limit(5).Documents(context.Background()).GetAll()

		var results []map[string]interface{}
		for _, doc := range docSnap {
			data := make(map[string]interface{})
			if err := doc.DataTo(&data); err != nil {
				log.Printf("failed to unmarshal document %s: %v", doc.Ref.ID, err)
				continue
			}
			results = append(results, data)
		}
	
		return results, nil
	}

	lastDocSnap, _ := db.Collection("blogMeta").Doc(lastDocId).Get(context.Background())

	docSnap, _ := db.Collection("blogMeta").Where("authorId", "==", userId).OrderBy("postDate", firestore.Desc).Limit(5).StartAfter(lastDocSnap).Documents(context.Background()).GetAll()

	var results []map[string]interface{}
	for _, doc := range docSnap {
		data := make(map[string]interface{})
		if err := doc.DataTo(&data); err != nil {
			log.Printf("failed to unmarshal document %s: %v", doc.Ref.ID, err)
			continue
		}
		results = append(results, data)
	}
	
	return results, nil
}