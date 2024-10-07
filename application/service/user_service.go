package service

/*自分用メモ:
 user情報を登録するユースケース
アプリケーション層では、ユースケースの処理を担当
バリデーションやビジネスロジックを行い、インフラ層にデータを委譲
*/
import (
	"context"

	"github.com/YuuHikida/GSC_backend_go/domain/model"
	"github.com/YuuHikida/GSC_backend_go/domain/repository"
	"github.com/YuuHikida/GSC_backend_go/infrastructure/external"
	"github.com/YuuHikida/GSC_backend_go/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(client *mongo.Client) UserService {
	return UserService{userRepository: persistence.NewMongoUserRepository(client)}
}

func (s *UserService) RegisterUser(userInfo model.UserInfo) (int, string) {
	// バリデーションチェック
	nRet, returnMsg := s.InputUserInfoValueCheckMain(userInfo)
	if nRet == 0 {
		return 0, returnMsg
	}

	err := s.userRepository.Save(&userInfo)
	if err != nil {
		return 0, "DB登録失敗"
	}
	return 1, "登録成功"
}

/*
		概要  : ユーザー登録情報のバリデーションチェック関数
		戻り値: nRet 0 = 異常, 1 = 正常
	        　　returnMsg 成功時、及び失敗時のメッセージ
*/
func (s *UserService) InputUserInfoValueCheckMain(stUserInfo model.UserInfo) (int, string) {

	// のチェック
	if nRet, returnMsg := checkName(stUserInfo.GitName); nRet == 0 {
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

func checkName(gitName string) (int, string) {

	// http.Getを使用してAccount文字列があるか確認
	exists, err := external.CheckGitHubAccount(gitName)
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

func (s *UserService) FindOneDocument(ctx context.Context, gitName string) (model.UserInfo, error) {
	return s.userRepository.FindOne(ctx, gitName)
}

func (s *UserService) FindAllDocuments(ctx context.Context) ([]model.UserInfo, error) {
	return s.userRepository.FindAll(ctx)
}
