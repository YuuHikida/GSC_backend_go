package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/YuuHikida/GSC_backend_go/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     *mongo.Client
	ctx        context.Context
	collection *mongo.Collection
)

// 接続情報を設定する関数
func SetDB(c *mongo.Client, context context.Context) {
	client = c
	ctx = context
	db := client.Database("gitInfoContributes")
	collection = db.Collection("user_info")
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	client, ctx, err := database.Initialize()
	if err != nil {
		log.Fatal("Database initialization failed: ", err)
	}

	fmt.Printf("Client: %v\n", client)
	fmt.Printf("Context: %v\n", ctx)

	db := client.Database("gitInfoContributes")
	collection := db.Collection("user_info")

	// デバッグ用にデータベース名とコレクション名を表示
	fmt.Printf("Database name: %s\n", db.Name())
	fmt.Printf("Collection name: %s\n", collection.Name())

	var result bson.M
	err_A := collection.FindOne(ctx, bson.M{"key": "value"}).Decode(&result)
	if err_A != nil {
		log.Printf("Error finding document: %v", err_A) // err_Aをログに出力する
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	} else {
		log.Printf("Document found: %+v", result)
	}

	log.Println("Found document:", result)
	w.Write([]byte("Document found: " + result["key"].(string)))
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
