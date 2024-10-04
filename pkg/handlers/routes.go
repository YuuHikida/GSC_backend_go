package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// ルーティング設定関数
func SetRoutes() http.Handler {
	router := mux.NewRouter()

	// ハンドラー設定
	// http.HnadleFuncよりrouterのほうがHTTPメソッドを指定可能
	router.HandleFunc("/", FindOne).Methods("GET") // HTTPメソッドを指定
	router.HandleFunc("/all", AllSelect).Methods("GET")

	// CORS設定 x
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Reactアプリのオリジンを許可
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// CORSミドルウェアを適用
	return c.Handler(router)
}
