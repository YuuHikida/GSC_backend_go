package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YuuHikida/GSC_backend_go/models"
	"github.com/YuuHikida/GSC_backend_go/pkg/validation"
	"github.com/YuuHikida/GSC_backend_go/services"
)

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
	/*-------------------
	 バリデーションチェック
	---------------------*/
	// リクエストボディをGoの構造体へデコード
	var body models.User_info
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	/*上記短縮系
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil
	 {...}*/

	nRet, returnMsg := validation.InputUserInfoValueCheckMain()

	// 戻り値判定 nRet(異常:->0,正常:->1)
	if nRet == 1 {
		http.Error(w, "入力値エラー:"+returnMsg, http.StatusBadRequest)
		return
	} else {
		return
	}
}
