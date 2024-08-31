package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDBクライアントの接続オプションを扱うパッケージ。
)

var client *mongo.Client // グローバル変数としてMongoDBクライアントを保持

func ConnectToMongoDB(uri string) *mongo.Client {

	// すでにクライアントが作成されている場合、そのクライアントを返す
	if client != nil {
		return client
	}

	// MongoDB 接続オプション設定
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB クライアントの作成
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err) // エラーが発生した場合、ログを出力してプログラムを終了
	}

	// 接続確認
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	return client
}

func GetURI() string {
	return "hello"
}
