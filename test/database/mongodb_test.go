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

var client *mongo.Client

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
			t.Errorf("returnURL :ERORR")
		}
		if got != want {
			t.Errorf("got :(%s), want:(%s)", got, want)
		}
	})

	t.Run("ConnectToMongoDB", func(t *testing.T) {
		moji := uri
		got, error := database.ConnectToMongoDB(client, moji)
		want := client
		if error != nil {
			t.Errorf("ConnectToMongoDB :ERORR")
		}
		if got != want {
			t.Errorf("got :(%v), want:(%v)", got, want)
		}
	})

	// t.Run("call DB and collection", func(t *testing.T) {
	// 	client, error := database.ConnectToMongoDB(uri, ctx)
	// 	got := database.CallDBAndcollection(client)
	// 	want := "test"
	// 	if error != nil {
	// 		t.Errorf("call DB and collection :ERORR")
	// 	}
	// 	if got != want {
	// 		t.Errorf("got :(%s), want:(%s)", got, want)
	// 	}
	// })

	t.Run("disconnectFunc", func(t *testing.T) {

	})
}
