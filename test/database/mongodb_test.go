package database

import (
	"os"
	"testing"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func TestMongoDB(t *testing.T) {
	// 準備
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		t.Fatal("MONGODB_URI is not set in .env file")
	}

	t.Run("call DB and collection", func(t *testing.T) {
		client := database.ConnectToMongoDB(uri)
		got := database.CallDBAndcollection(client)
		want := "test"

		if got != want {
			t.Errorf("got :(%s), want:(%s)", got, want)
		}
	})
	t.Run("connect", func(t *testing.T) {
		element := uri
		got := database.ConnectToMongoDB(element)
		want := client

		if got != want {
			t.Errorf("got :(%v), want:(%v)", got, want)
		}
	})

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

	t.Run("disconnectFunc", func(t *testing.T) {

	})
}
