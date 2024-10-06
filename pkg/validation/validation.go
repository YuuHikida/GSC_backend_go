package validation

import (
	"github.com/YuuHikida/GSC_backend_go/models"
	"github.com/YuuHikida/GSC_backend_go/pkg/service"
)

/*
		概要  : ユーザー登録情報のバリデーションチェック関数
		戻り値: nRet 0 = 異常, 1 = 正常
	        　　returnMsg 成功時、及び失敗時のメッセージ
*/
func InputUserInfoValueCheckMain(stUserInfo models.User_info) (int, string) {

	// のチェック
	if nRet, returnMsg := check(stUserInfo.GitName); nRet == 0 {
		return nRet, returnMsg
	}

	// Mailのチェック
	if nRet, returnMsg := checkMail(stUserInfo.Mail); nRet == 0 {
		return nRet, returnMsg
	}

	// 時間のチェック
	if nRet, returnMsg := checkTime(stUserInfo.Time); nRet == 0 {
		return nRet, returnMsg
	}

	// 全ての値が正常
	nRet := 1
	return nRet, "全ての入力が正常"
}

func check(gitName string) (int, string) {

	// http.Getを使用してAccount文字列があるか確認
	exists, err := service.CheackGitHubAccout(gitName)
	if err != nil {
		return 0, "http.Get取得エラー"
	}
	if !exists { //←false
		return 0, "このGitIDは存在しません"
	}

	return 1, "Git Name 正常"
}

func checkMail(mail string) (int, string) {
	// チェック処理
	return 1, "Mail 正常"
}

func checkTime(time string) (int, string) {
	// チェック処理
	return 1, "Time 正常"
}
