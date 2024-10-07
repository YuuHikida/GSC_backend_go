package model

/*
	前提知識
	"構造体名" & "メンバ変数"の先頭1文字目は大文字か、小文字かで挙動が変わる
	大文字:パッケージ外からアクセス出来る
	小文字:パッケージ内からアクセス

	'bson:xxxは DBのドキュメント名,jsonはgo構造体とjsonのデータをパース
*/
// ユーザー登録情報
type UserInfo struct {
	GitName string `bson:"git_name" json:"git_name"`
	Mail    string `bson:"mail" json:"mail"`
	Time    string `bson:"time" json:"time"`
}
