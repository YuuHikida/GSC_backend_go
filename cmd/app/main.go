package main

import (
	"fmt"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
)

func main() {

	fmt.Println("-------------- ! Process Start ! --------------")
	// MongoDB call
	uri := database.GetURI()
	client := database.ConnectToMongoDB(uri)

	ans := database.CallDBAndcollection(client)

	fmt.Println(ans)
	fmt.Println("---------------- ! Process End ! ----------------")
}
