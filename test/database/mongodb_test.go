package database

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestMongoDB(t *testing.T) {

	t.Run("connect", func(t *testing.T) {
		err := godotenv.Load()
		if err != nil {
			t.Fatalf("Error loading .env file")
		}

		element := "URL"
		got := ConnectToMongoDB(element)
		want := "HIKIDA"

		if got != want {
			t.Errorf("got :(%g), want:(%g)", got, want)
		}
	})
}
