package api

import (
	"context"
	"encoding/json"
	"log"

	"net/http"

	"github.com/YuuHikida/GSC_backend_go/application/service"
	"github.com/YuuHikida/GSC_backend_go/domain/model"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// 一件のドキュメントを取得してJSONで返す
func (h *UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	gitName := "TANAKA" // クエリパラメータやリクエストから取得しても良い

	result, err := h.userService.FindOneDocument(context.Background(), gitName)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		http.Error(w, "Database error_findOne", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// 全件取得してJSONで返す
func (h *UserHandler) AllSelect(w http.ResponseWriter, r *http.Request) {
	results, err := h.userService.FindAllDocuments(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// ユーザー情報を登録する
func (h *UserHandler) RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	var body model.UserInfo

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// バリデーションチェック及び登録
	//戻り値は成功or失敗値とメッセ
	nRet, returnMsg := h.userService.RegisterUser(body)

	// HTTPレスポンスのヘッダーを設定
	w.Header().Set("Content-Type", "text/plain") // w.Header().Set("Content-Type", "application/json") <- これはJsonを返す
	w.Write([]byte(returnMsg))

	// 一旦console.log出力
	//いずれはメッセージを登録画面に返す
	log.Printf("nRet:(%d),retunrMSG:(%s)", nRet, returnMsg)

	//w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) editUserInfo(w http.ResponseWriter, r *http.Request) {
	// ここで値を編集

}
