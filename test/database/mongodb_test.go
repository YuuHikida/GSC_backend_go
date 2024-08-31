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

	t.Run("connect", func(t *testing.T) {
		element := "URL"
		got := database.ConnectToMongoDB(element)
		want := "HIKIDA"

		if got != want {
			t.Errorf("got :(%q), want:(%q)", got, want)
		}
	})
}
