package external

// 概要:# ビジネスロジックをサービスとして定義
import (
	"fmt"
	"net/http"
)

/*
概要  : GitHubのアカウントID文字列が存在するかの確認
*/
func CheckGitHubAccount(userName string) (bool, error) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", userName)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// ステータスコードが200だった場合は成功
	return resp.StatusCode == http.StatusOK, nil

}
