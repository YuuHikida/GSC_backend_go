package model

/*
	前提知識
	"構造体名" & "メンバ変数"の先頭1文字目は大文字か、小文字かで挙動が変わる
	大文字:パッケージ外からアクセス出来る
	小文字:パッケージ内からアクセス
*/
// ユーザー登録情報
type UserInfo struct {
	GitName string
	Mail    string
	Time    string
}
