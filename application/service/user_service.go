package service

/*自分用メモ:
 user情報を登録するユースケース
アプリケーション層では、ユースケースの処理を担当
バリデーションやビジネスロジックを行い、インフラ層にデータを委譲
*/
import (
	"context"
	"errors"

	"github.com/YuuHikida/GSC_backend_go/domain/model"
	"github.com/YuuHikida/GSC_backend_go/domain/repository"
	"github.com/YuuHikida/GSC_backend_go/infrastructure/external"
	"github.com/YuuHikida/GSC_backend_go/infrastructure/persistence"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return UserService{userRepository: persistence.NewMongoUserRepository()}
}

func (s *UserService) RegisterUser(userInfo model.UserInfo) error {
	if err := s.validateUser(userInfo); err != nil {
		return err
	}

	return s.userRepository.Save(userInfo)
}

func (s *UserService) validateUser(userInfo model.UserInfo) error {
	if userInfo.GitName == "" || userInfo.Mail == "" {
		return errors.New("invalid user info")
	}
	return nil
}

/*
		概要  : ユーザー登録情報のバリデーションチェック関数
		戻り値: nRet 0 = 異常, 1 = 正常
	        　　returnMsg 成功時、及び失敗時のメッセージ
*/
func (s *UserService) InputUserInfoValueCheckMain(stUserInfo model.UserInfo) (int, string) {

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
