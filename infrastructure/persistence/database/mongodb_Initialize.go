package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDBクライアントの接続オプションを扱うパッケージ。
)

/*
  --- MEMO --
	[context.Context の役割]:
	・context.Context は、キャンセルやタイムアウトの制御、
	リクエスト間の値の伝播、または異なる操作間での共有のために使用される

	[コンテキスト（Context）とは？]
	・コンテキストは、Go言語でキャンセル信号やデッドライン（タイムアウト）、
		その他のリクエストに関連する情報を管理するための構造体です。
			一般的に、複数のゴルーチン（並行処理のスレッド）での処理やリクエストのキャンセルを
				一元的に管理するために使用されます。

	・コンテキストを使うことで、リクエストのスコープ内で、
		例えば、タイムアウトやキャンセルなどの状態を共有することができます。

*/

// var client *mongo.Client // グローバル変数としてMongoDBクライアントを保持

func Initialize() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	uri, err := GetURI()
	if err != nil {
		log.Fatalf("Failed to get URI: %v", err)
	}

	client, err := ConnectToMongoDB(ctx, uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return client, ctx, cancel, err
}

// mongoDB Atlas 接続文字列取得
func GetURI() (string, error) {

	// Docker使わずにローカルで起動するなら以下
	// err := godotenv.Load("../../.env")
	// err := godotenv.Load("/app/.env")
	// if err != nil {
	// 	// log.Fatalfはエラーメッセをフォーマットして出力し、プログラムを終了する
	// 	log.Fatalf("Error loading .env file :(%v)", err)
	// 	// %v：変数のデフォルトの形式で表示します。構造体など複雑なデータのデフォルト表示に使います
	// }

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("MONGODB_URI is not set in .env file")
	}

	return uri, nil
}

// MongoDBクライアントを取得する関数
func ConnectToMongoDB(ctx context.Context, uri string) (*mongo.Client, error) {

	// すでにクライアントが作成されている場合、そのクライアントを返す
	// if client != nil {
	// 	return client
	// }

	// MongoDB 接続オプション設定
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB クライアントの作成
	// Connectにより作成されたインスタンスは破棄されない限り、繋がり続ける(破棄)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// log.Fatalf("Error creating MongoDB client: %w", err) // エラーが発生した場合、ログを出力してプログラムを終了
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	// 接続確認
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}
	return client, nil
}

// // 指定されたクライアントを使用して、コレクションを取得する関数
// func GetCollection(client *mongo.Client, dbName, collectionName string) (*mongo.Collection, error) {
// 	// NULLチェック
// 	if client == nil {
// 		return nil, fmt.Errorf("client is nil")
// 	}

// 	collection := client.Database(dbName).Collection(collectionName)
// 	return collection, nil
// }

// // クライアントの接続を閉じる関数
// func DisconnectClient(ctx context.Context, client *mongo.Client) error {
// 	if err := client.Disconnect(ctx); err != nil {
// 		return fmt.Errorf("error disconnecting from MongoDB: %w", err)
// 	}
// 	return nil
// }
