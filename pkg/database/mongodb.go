package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
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

// mongoDB Atlas 接続文字列取得
func GetURI() string {

	err := godotenv.Load("../../.env")
	if err != nil {
		// log.Fatalfはエラーメッセをフォーマットして出力し、プログラムを終了する
		log.Fatalf("Error loading .env file :(%v)", err)
		// %v：変数のデフォルトの形式で表示します。構造体など複雑なデータのデフォルト表示に使います
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("MONGODB_URI is not set in .env file")
	}

	return uri
}
