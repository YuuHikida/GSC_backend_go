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

	"github.com/YuuHikida/GSC_backend_go/pkg/api"
	"github.com/YuuHikida/GSC_backend_go/pkg/database"
	"github.com/YuuHikida/GSC_backend_go/services"
)

// 書き残し:2024-09-01 net/httpの書き方とイベントハンドラーについて学ぶたびに出るのでいったん作業終了
// 現状8080でルートアクセスしてもDBの値取得できず
func main() {
	fmt.Println("-- Start Program --")

	//　初期設定(DBの初期化)
	client, ctx, cancel, err := database.Initialize()
	if err != nil {
		log.Fatal("Database initialization failed: ", err)
	}

	// コンテキストキャンセルをサーバーが終了するタイミングで実行
	defer cancel()

	// サーバーが終了する際に接続を切断する
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("Error disconnecting from the database:", err)
		}
	}()

	// サービス層の初期化
	services.Initialize(client)

	// ハンドラー設定
	http.HandleFunc("/", api.FindOne)
	http.HandleFunc("/all", api.AllSelect)

	// サーバー起動
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)

	fmt.Println("-- END Program --")
}
