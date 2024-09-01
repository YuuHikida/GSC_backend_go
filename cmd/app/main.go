package main

import (
	"fmt"
	"log"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
)

func main() {

	fmt.Println("-------------- ! Process Start ! --------------")
	// MongoDB call
	uri, err := database.GetURI()
	if err != nil {
		log.Fatalf("Failed to get URI: %v", err)
	}
	client := database.ConnectToMongoDB(uri)

	ans := database.CallDBAndcollection(client)

	fmt.Println(ans)

	fmt.Println("---------------- ! Process End ! ----------------")
}
