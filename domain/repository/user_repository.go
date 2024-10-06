package repository

import (
	"context"

	"github.com/YuuHikida/GSC_backend_go/domain/model"
)

type UserRepository interface {
	Save(user *model.UserInfo) error // ポインタ型で受け取る
	FindOne(ctx context.Context, gitName string) (model.UserInfo, error)
	FindAll(ctx context.Context) ([]model.UserInfo, error)
}
