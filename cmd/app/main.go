package main

import (
	"fmt"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
)

func main() {

	// MongoDB call
	uri := database.GetURI()
	client := database.ConnectToMongoDB(uri)

	fmt.Println("Hello,World!")
}
