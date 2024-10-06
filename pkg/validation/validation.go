package validation

import "github.com/YuuHikida/GSC_backend_go/models"

/*
	概要  : ユーザー登録情報のバリデーションチェック関数
	戻り値: nRet 0 = 異常, 1 = 正常
        　　returnMsg 成功時、及び失敗時のメッセージ
*/
func InputUserInfoValueCheckMain(stUserInfo models.User_info) (int, string) {
	// チェックする値の初期化
	var (
		GitName string = stUserInfo.GitName
		Mail    string = stUserInfo.Mail
		Time    string = stUserInfo.Time
	)
	// GitNameのチェック
	if nRet, returnMsg := checkGitName(GitName); nRet == 0 {
		return nRet, returnMsg
	}

	// Mailのチェック
	if nRet, returnMsg := checkMail(Mail); nRet == 0 {
		return nRet, returnMsg
	}

	// 時間のチェック
	if nRet, returnMsg := checkTime(Time); nRet == 0 {
		return nRet, returnMsg
	}

	// 全ての値が正常
	nRet := 1
	return nRet, "全ての入力が正常"
}

func checkGitName(GitName string) (int, string) {
	// チェック処理
	return 1, "Git Name 正常"
}

func checkMail(Mail string) (int, string) {
	// チェック処理
	return 1, "Mail 正常"
}

func checkTime(Time string) (int, string) {
	// チェック処理
	return 1, "Time 正常"
}
