package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func TestMongoDB(t *testing.T) {
	// 準備
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		t.Fatal("MONGODB_URI is not set in .env file")
	}

	t.Run("connect", func(t *testing.T) {

		element := "URL"
		got := ConnectToMongoDB(element)
		want := "HIKIDA"

		if got != want {
			t.Errorf("got :(%g), want:(%g)", got, want)
		}
	})
}
