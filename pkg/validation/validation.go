package validation

/*
	概要  : ユーザー登録情報のバリデーションチェック関数
	戻り値: nRet 0 = 異常, 1 = 正常
        　　returnMsg 成功時、及び失敗時のメッセージ
*/
func InputUserInfoValueCheckMain() (int, string) {
	// GitNameのチェック
	if nRet, returnMsg := checkGitName(); nRet == 0 {
		return nRet, returnMsg
	}

	// Mailのチェック
	if nRet, returnMsg := checkMail(); nRet == 0 {
		return nRet, returnMsg
	}

	// 時間のチェック
	if nRet, returnMsg := checkTime(); nRet == 0 {
		return nRet, returnMsg
	}

	return 1, "全ての入力が正常"
}

func checkGitName() (int, string) {

	return nRet, returnMsg
}
func checkMail() (int, string) {

}
func checkTime() (int, string) {

}
