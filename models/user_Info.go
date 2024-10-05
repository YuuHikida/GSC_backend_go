package models

/*
	前提知識
	メンバ変数の先頭1文字目は大文字か、小文字かで挙動が変わる
	大文字:パッケージ外からアクセス出来る
	小文字:パッケージ内からアクセス
*/
// ユーザー登録情報
type user_info struct {
	GitName string
	Mail    string
	time    string
}
