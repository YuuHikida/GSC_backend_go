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
	//　初期設定
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

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func main() {
// 	// 設定とデータベースの初期化

// 	db, ctx, err := database.Initialize()
// 	if err != nil {
// 		log.Fatal("Database initialization failed: ", err)
// 	}

// 	// ルーティングの設定
// 	http.HandleFunc("/", api.HandleRoot(db, ctx)) // ルートハンドラーを登録

// 	// サーバ起動
// 	log.Println("Server starting on port 8080...")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

//--------------------------------------------------------------------------------------
// func main() {
// 	fmt.Println("-------------- ! Process Start ! --------------")

// 	// コンテキスト作成
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// 接続文字列取得
// 	uri, err := database.GetURI()
// 	if err != nil {
// 		log.Fatalf("Failed to get URI: %v", err)
// 	}

// 	// MongoDBクライアントを接続
// 	client, err := database.ConnectToMongoDB(ctx, uri)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer func() {
// 		database.DisconnectClient(ctx, client)
// 	}()

// 	// コレクションを取得
// 	collection, err := database.GetCollection(client, "gitInfoContributes", "user_info")
// 	if err != nil {
// 		log.Fatalf("Failed to get collection: %v", err)
// 	}

// 	// ここでコレクションを使った操作を行う...
// 	// 例: fmt.Println(collection)

// 	// ドキュメント取得
// 	count, err := collection.CountDocuments(ctx, bson.D{})
// 	if err != nil {
// 		log.Fatalf("Failed to count documents: %v", err)
// 	}
// 	fmt.Printf("Documents count: %d\n", count)

// 	/*
// 		bson.D は、Goの構造体である bson.E のスライスで構成されている
// 		bson.E というのは、キーと値を保持するための構造体で、以下のように定義

// 		type E struct {
// 			Key   string
// 			Value interface{}
// 		}

// 	*/
// 	// ドキュメント取得(1件)
// 	var result bson.M
// 	err_result := collection.FindOne(ctx, bson.D{{Key: "git_name", Value: "TANAKA"}}).Decode(&result)
// 	if err_result != nil {
// 		log.Fatalf("Failed to find document: %v", err_result)
// 	}
// 	fmt.Printf("Retrieved document: %v\n", result)

// 	fmt.Println("---------------- ! Process End ! ----------------")
// }
