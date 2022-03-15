package repository

import (
	"22dojo-online/pkg/domain/entity"
)

type UserRepository interface {
	SelectUserByPrimaryKey(string) (*entity.User, error)
	SelectUserByAuthToken(string) (*entity.User, error)
	InsertUser(user *entity.User) error
	UpdateUserByPrimaryKey(user *entity.User) error
}
