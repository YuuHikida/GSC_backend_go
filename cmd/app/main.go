package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/YuuHikida/GSC_backend_go/pkg/database" // databaseパッケージのインポート
)

func main() {
	fmt.Println("-------------- ! Process Start ! --------------")

	// コンテキスト作成
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 接続文字列取得
	uri, err := database.GetURI()
	if err != nil {
		log.Fatalf("Failed to get URI: %v", err)
	}

	// MongoDBクライアントを接続
	client, err := database.ConnectToMongoDB(ctx, uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := database.DisconnectClient(ctx, client); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	ans := database.CallDBAndcollection(client)

	fmt.Println(ans)

	fmt.Println("---------------- ! Process End ! ----------------")
}
