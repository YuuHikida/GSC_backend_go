package main

/*
	log: ロギングを行うための標準ライブラリ。エラーが発生した際にログ出力できるように使用する。
	net/http: HTTPサーバーとクライアント機能を提供する標準ライブラリ。サーバーを立てたり、HTTPリクエストを処理したりするのに必要。
	yourapp/pkg/api: 独自に定義したAPIハンドラーが含まれるパッケージ。具体的なリクエスト処理を行う関数をここから利用する。
	yourapp/pkg/config: 設定を管理するパッケージ。設定ファイルや環境変数から設定を読み込む関数が含まれる。
	yourapp/pkg/database: データベースとの接続を管理するパッケージ。MongoDBとの接続を確立するためのロジックが含まれる。
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/YuuHikida/GSC_backend_go/infrastructure/persistence"
	"github.com/YuuHikida/GSC_backend_go/infrastructure/persistence/database"
	"github.com/YuuHikida/GSC_backend_go/interfaces/api"
)

func main() {
	fmt.Println("-- Start Program --")

	//　DBの初期化
	client, ctx, cancel, err := database.Initialize()
	if err != nil {
		log.Fatal("Database initialization failed: ", err)
	} else {
		log.Printf("Succesfully connected")
	}
	// コンテキストキャンセルをサーバーが終了するタイミングで実行
	defer cancel()

	// サーバーが終了する際に接続を切断する
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("Error disconnecting from the database:", err)
		}
	}()

	// ルーターを作成
	handler := api.SetRoutes(client)

	// リポジトリ初期化（必要なら後で利用する）
	userRepo := persistence.NewMongoUserRepository(client)
	if userRepo == nil {
		log.Printf("リポジトリ初期化失敗")
	}
	// サーバー起動
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler)) // エラーハンドリングを追加

	fmt.Println("-- END Program --")
}
