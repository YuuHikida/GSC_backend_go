package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YuuHikida/GSC_backend_go/services"
)

/*
	bson.Dとbson.M違いはキーと値ぺの順番を重要

bson.D... 構造体の型: []bson.E（bson.Eは、キーと値のペアを表す構造体）
bson.M... 構造体の型: map[string]interface{}（Goのマップ構造に似たもので、キーと値のペア）
*/

// var (
// 	client     *mongo.Client
// 	ctx        context.Context
// 	collection *mongo.Collection
// )

// 一件のドキュメントを取得してJSONで返す
func FindOne(w http.ResponseWriter, r *http.Request) {
	gitName := "TANAKA" // クエリパラメータやリクエストから取得しても良い

	result, err := services.FindOneDocument(context.Background(), gitName)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// JSONレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// 全件取得してJSONで返す
func AllSelect(w http.ResponseWriter, r *http.Request) {
	results, err := services.FindAllDocuments(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// JSONレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// user情報を登録する
func RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	// バリデーションチェック

}
