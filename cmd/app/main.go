package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("-------------- ! Process Start ! --------------")

	// コンテキスト作成
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 接続文字列取得
	uri, err := database.GetURI()
	if err != nil {
		log.Fatalf("Failed to get URI: %v", err)
	}

	// MongoDBクライアントを接続
	client, err := database.ConnectToMongoDB(ctx, uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		database.DisconnectClient(ctx, client)
	}()

	// コレクションを取得
	collection, err := database.GetCollection(client, "gitInfoContributes", "user_info")
	if err != nil {
		log.Fatalf("Failed to get collection: %v", err)
	}

	// ここでコレクションを使った操作を行う...
	// 例: fmt.Println(collection)

	// ドキュメント取得
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}
	fmt.Printf("Documents count: %d\n", count)

	/*
		bson.D は、Goの構造体である bson.E のスライスで構成されている
		bson.E というのは、キーと値を保持するための構造体で、以下のように定義

		type E struct {
			Key   string
			Value interface{}
		}

	*/
	// ドキュメント取得(1件)
	var result bson.M
	err_result := collection.FindOne(ctx, bson.D{{Key: "git_name", Value: "TANAKA"}}).Decode(&result)
	if err_result != nil {
		log.Fatalf("Failed to find document: %v", err_result)
	}
	fmt.Printf("Retrieved document: %v\n", result)

	fmt.Println("---------------- ! Process End ! ----------------")
}
