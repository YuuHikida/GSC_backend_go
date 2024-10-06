package repository

import "github.com/YuuHikida/GSC_backend_go/domain/model"

type UserRepository interface {
	Save(user model.User_info) error
}
