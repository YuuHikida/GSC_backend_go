package api

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleRoot(client *mongo.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("gitInfoContributes")
		collection := database.Collection("user_info")

		var result bson.M
		err := collection.FindOne(ctx, bson.M{"key": "value"}).Decode(&result)
		if err != nil {
			log.Printf("Error finding document: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		log.Println("Found document:", result)
		w.Write([]byte("Document found: " + result["key"].(string)))
	}
}

// func HandleRoot(client *mongo.Client, ctx context.Context) http.HandlerFunc {
// 	// データベースを選択
// 	database := client.Database("gitInfoContributes")
// 	// コレクションを選択
// 	collection := database.Collection("user_info")

// 	// 検索クエリ
// 	var result bson.M
// 	// err := collection.FindOne(context.TODO(), bson.M{"key": "value"}).Decode(&result)
// 	err := collection.FindOne(ctx, bson.M{"key": "value"}).Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Found document:", result)
// }
