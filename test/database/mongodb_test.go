package database

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoDB(t *testing.T) {
	//------------- 準備 -----------------------
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		t.Fatal("MONGODB_URI is not set in .env file")
	}
	// コンテキスト作成
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//MongoDBクライアント
	clientOptions := options.Client().ApplyURI(uri) // MongoDB 接続オプション設定
	client, err := mongo.Connect(ctx, clientOptions)

	//---------------------------------------------
	t.Run("returnURI", func(t *testing.T) {
		got, error := database.GetURI()
		want := uri
		if error != nil {
			t.Errorf("returnURI :ERROR")
		}
		if got != want {
			t.Errorf("got :(%s), want:(%s)", got, want)
		}
	})

	// t.Run("ConnectToMongoDB", func(t *testing.T) {
	// 	moji := uri
	// 	got, err := database.ConnectToMongoDB(ctx, moji) // 修正: ctx を第1引数に渡す
	// 	want := client
	// 	if err != nil {
	// 		t.Errorf("ConnectToMongoDB :ERROR")
	// 	}
	// 	if got != want {
	// 		t.Errorf("got :(%v), want:(%v)", got, want)
	// 	}
	// })

	t.Run("disconnectFunc", func(t *testing.T) {
		// ここでDisconnectのテストができるよ
		err := database.DisconnectClient(ctx, client)
		if err != nil {
			t.Errorf("Failed to disconnect MongoDB client: %v", err)
		}
	})
}
