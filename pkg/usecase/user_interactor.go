package usecase

import (
	"fmt"

	"22dojo-online/pkg/domain/entity"
	"22dojo-online/pkg/domain/service"
)

type UserInteractorInterface interface {
	SelectUserByPrimaryKey(string) (*entity.User, error)
	SelectUserByAuthToken(string) (*entity.User, error)
	InsertUser(user *entity.User) error
	UpdateUserByPrimaryKey(user *entity.User) error
}

type UserInteractor struct {
	UserService service.UserServiceInterface
}

func NewUserInteractor(userService service.UserServiceInterface) UserInteractorInterface {
	return &UserInteractor{
		UserService: userService,
	}
}

func (ui *UserInteractor) SelectUserByPrimaryKey(userID string) (*entity.User, error) {
	user, err := ui.UserService.SelectUserByPrimaryKey(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query a user by primary key. err = %w", err)
	}
	return user, nil
}

func (ui *UserInteractor) SelectUserByAuthToken(authToken string) (*entity.User, error) {
	user, err := ui.UserService.SelectUserByAuthToken(authToken)
	if err != nil {
		return nil, fmt.Errorf("failed to query a user by auth token. err = %w", err)
	}
	return user, nil
}

// request paramのnameだけを受け取り
func (ui *UserInteractor) InsertUser(user *entity.User) error {
	if err := ui.UserService.InsertUser(user); err != nil {
		return fmt.Errorf("failed to insert user. err = %w", err)
	}
	return nil
}

func (ui *UserInteractor) UpdateUserByPrimaryKey(user *entity.User) error {
	if err := ui.UserService.UpdateUserByPrimaryKey(user); err != nil {
		return fmt.Errorf("failed to update user. err = %w", err)
	}
	return nil
}
