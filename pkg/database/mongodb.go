package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDBクライアントの接続オプションを扱うパッケージ。
)

/*
  --- MEMO --
	context.Context の役割:
	・context.Context は、キャンセルやタイムアウトの制御、
	リクエスト間の値の伝播、または異なる操作間での共有のために使用される


*/

var client *mongo.Client // グローバル変数としてMongoDBクライアントを保持

// mongoDB Atlas 接続文字列取得
func GetURI() (string, error) {

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

	return uri, nil
}

func CallDBAndcollection(client *mongo.Client) string {
	// NULLチェック
	if client == nil {
		log.Fatalf("Database connect error : (%v)", client)
	}

	db := client.Database("gitInfoContributes")
	collection := db.Collection("user_info")

	fmt.Println("%s", collection)
	return "hellohelllo"

}

func ConnectToMongoDB(uri string) *mongo.Client {

	// すでにクライアントが作成されている場合、そのクライアントを返す
	// if client != nil {
	// 	return client
	// }

	// MongoDB 接続オプション設定
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB クライアントの作成
	// Connectにより作成されたインスタンスは破棄されない限り、繋がり続ける(破棄)
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

// クライアントの接続を閉じる関数
func DisconnectClient(ctx context.Context, client *mongo.Client) error {
	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %w", err)
	}
	return nil
}
