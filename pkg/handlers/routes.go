package handlers

import (
	"github.com/YuuHikida/GSC_backend_go/pkg/handlers"
	"github.com/gorilla/mux"
)

// ルーティング設定関数
func SetRoutes() *mux.Router {
	router = mux.NewRouter()

	// ハンドラー設定
	// http.HnadleFuncよりrouterのほうがHTTPメソッドを指定可能
	router.HandleFunc("/", handlers.FindOne).Methods("GET") // HTTPメソッドを指定
	router.HandleFunc("/all", handlers.AllSelect).Methods("GET")

	return router
}
