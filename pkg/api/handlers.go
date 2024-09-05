package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

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

// DBの要素を一件取得
func FindOne(w http.ResponseWriter, r *http.Request) {

	if client == nil || collection == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}
	// デバッグ用にデータベース名とコレクション名を表示
	fmt.Printf("Database name: %s\n", collection.Database().Name())
	fmt.Printf("Collection name: %s\n", collection.Name())

	// MongoDB クライアントが接続されているか確認
	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("MongoDB client is disconnected: %v", err)
		http.Error(w, "Database connection lost", http.StatusInternalServerError)
		return
	}

	/* bson.Dとbson.M違いはキーと値ぺの順番を重要
	bson.D... 構造体の型: []bson.E（bson.Eは、キーと値のペアを表す構造体）
	bson.M... 構造体の型: map[string]interface{}（Goのマップ構造に似たもので、キーと値のペア）
	*/
	var result bson.M

	err_A := collection.FindOne(ctx, bson.M{"git_name": "TANAKA"}).Decode(&result)
	if err_A != nil {
		log.Printf("Error finding document: %v", err_A) // err_Aをログに出力する
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	log.Println("Found document:", result)

	if gitName, ok := result["git_name"].(string); ok {
		w.Write([]byte("Document found: " + gitName))
	} else {
		log.Println("Key 'git_name' is missing or not a string")
		http.Error(w, "Key 'git_name' not found or invalid", http.StatusInternalServerError)
		return
	}
}

// DBのcollectionを全件取得
func AllSelect(w http.ResponseWriter, r *http.Request) {

	if client == nil || collection == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}
	// デバッグ用にデータベース名とコレクション名を表示
	fmt.Printf("Database name: %s\n", collection.Database().Name())
	fmt.Printf("Collection name: %s\n", collection.Name())

	// MongoDB クライアントが接続されているか確認
	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("MongoDB client is disconnected: %v", err)
		http.Error(w, "Database connection lost", http.StatusInternalServerError)
		return
	}

	// 全レコードを取得するクエリ
	filter := bson.D{}

	cursor, err_A := collection.Find(ctx, filter)
	if err_A != nil {
		log.Printf("Error finding document: %v", err_A) // err_Aをログに出力する
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	log.Println("Found document:", result)

	if gitName, ok := result["git_name"].(string); ok {
		w.Write([]byte("Document found: " + gitName))
	} else {
		log.Println("Key 'git_name' is missing or not a string")
		http.Error(w, "Key 'git_name' not found or invalid", http.StatusInternalServerError)
		return
	}
}

// func AllSelect(client *mongo.Client, ctx context.Context) http.HandlerFunc {
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
