package mock_service

import (
	"testing"

	entity "22dojo-online/pkg/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// test
func TestSelectUserByPrimaryKeyService(t *testing.T) {
	table := []struct {
		testName string
		id       string
		user     entity.User
		err      error
	}{
		{
			"FIRST TEST CASE: TestSelectUserByPrimaryKeyService",
			"78164dcf-6b7c-45e4-862a-2a0f6735a449",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 100,
				Coin:      10000,
			},
			nil,
		},
		{
			"SECOND TEST CASE: TestSelectUserByPrimaryKeyService",
			"829bdb53-f322-40bb-9327-63ab00536cd3",
			entity.User{
				ID:        "829bdb53-f322-40bb-9327-63ab00536cd3",
				AuthToken: "4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
				Name:      "bruh",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
		{
			"THIRD TEST CASE: TestSelectUserByPrimaryKeyService",
			"909b42f8-cce9-4d02-bce4-e7c7e28df550",
			entity.User{
				ID:        "909b42f8-cce9-4d02-bce4-e7c7e28df550",
				AuthToken: "bb68df68-964e-4f27-a225-0cbafdd6ce9f",
				Name:      "9",
				HighScore: 70,
				Coin:      0,
			},
			nil,
		},
	}
	fservice := gomock.NewController(t)
	m := NewMockUserServiceInterface(fservice)

	for _, tt := range table {
		m.EXPECT().SelectUserByPrimaryKey(tt.id).Return(&tt.user, nil)
		user, err := m.SelectUserByPrimaryKey(tt.id)
		assert.Equal(t, user, &tt.user)
		assert.Equal(t, nil, err)
	}
}

func TestSelectUserByAuthTokenService(t *testing.T) {
	table := []struct {
		testName  string
		authToken string
		user      entity.User
		err       error
	}{
		{
			"FIRST TEST CASE: TestSelectUserByPrimaryKeyService",
			"b187b9e0-08e6-42dd-a9b3-a900b137983c",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 100,
				Coin:      10000,
			},
			nil,
		},
		{
			"SECOND TEST CASE: TestSelectUserByPrimaryKeyService",
			"4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
			entity.User{
				ID:        "829bdb53-f322-40bb-9327-63ab00536cd3",
				AuthToken: "4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
				Name:      "bruh",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
		{
			"THIRD TEST CASE: TestSelectUserByPrimaryKeyService",
			"bb68df68-964e-4f27-a225-0cbafdd6ce9f",
			entity.User{
				ID:        "909b42f8-cce9-4d02-bce4-e7c7e28df550",
				AuthToken: "bb68df68-964e-4f27-a225-0cbafdd6ce9f",
				Name:      "9",
				HighScore: 70,
				Coin:      0,
			},
			nil,
		},
	}
	fservice := gomock.NewController(t)
	m := NewMockUserServiceInterface(fservice)

	for _, tt := range table {
		m.EXPECT().SelectUserByAuthToken(tt.authToken).Return(&tt.user, nil)
		user, err := m.SelectUserByAuthToken(tt.authToken)
		assert.Equal(t, user, &tt.user)
		assert.Equal(t, nil, err)
	}
}

func TestInsertUserService(t *testing.T) {
	table := []struct {
		testName  string
		authToken string
		user      entity.User
		err       error
	}{
		{
			"FIRST TEST CASE: TestSelectUserByPrimaryKeyService",
			"b187b9e0-08e6-42dd-a9b3-a900b137983c",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 100,
				Coin:      10000,
			},
			nil,
		},
		{
			"SECOND TEST CASE: TestSelectUserByPrimaryKeyService",
			"4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
			entity.User{
				ID:        "829bdb53-f322-40bb-9327-63ab00536cd3",
				AuthToken: "4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
				Name:      "bruh",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
		{
			"THIRD TEST CASE: TestSelectUserByPrimaryKeyService",
			"bb68df68-964e-4f27-a225-0cbafdd6ce9f",
			entity.User{
				ID:        "909b42f8-cce9-4d02-bce4-e7c7e28df550",
				AuthToken: "bb68df68-964e-4f27-a225-0cbafdd6ce9f",
				Name:      "9",
				HighScore: 70,
				Coin:      0,
			},
			nil,
		},
	}
	fservice := gomock.NewController(t)
	m := NewMockUserServiceInterface(fservice)

	for _, tt := range table {
		m.EXPECT().InsertUser(&tt.user).Return(nil)
		err := m.InsertUser(&tt.user)
		assert.Equal(t, nil, err)
	}
}

func TestUpdateUserByPrimaryKey(t *testing.T) {
	table := []struct {
		testName  string
		authToken string
		user      entity.User
		err       error
	}{
		{
			"FIRST TEST CASE: TestSelectUserByPrimaryKeyService",
			"b187b9e0-08e6-42dd-a9b3-a900b137983c",
			entity.User{
				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
				Name:      "whatt",
				HighScore: 100,
				Coin:      10000,
			},
			nil,
		},
		{
			"SECOND TEST CASE: TestSelectUserByPrimaryKeyService",
			"4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
			entity.User{
				ID:        "829bdb53-f322-40bb-9327-63ab00536cd3",
				AuthToken: "4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
				Name:      "bruh",
				HighScore: 0,
				Coin:      0,
			},
			nil,
		},
		{
			"THIRD TEST CASE: TestSelectUserByPrimaryKeyService",
			"bb68df68-964e-4f27-a225-0cbafdd6ce9f",
			entity.User{
				ID:        "909b42f8-cce9-4d02-bce4-e7c7e28df550",
				AuthToken: "bb68df68-964e-4f27-a225-0cbafdd6ce9f",
				Name:      "9",
				HighScore: 70,
				Coin:      0,
			},
			nil,
		},
	}
	fservice := gomock.NewController(t)
	m := NewMockUserServiceInterface(fservice)

	for _, tt := range table {
		m.EXPECT().UpdateUserByPrimaryKey(&tt.user).Return(nil)
		err := m.UpdateUserByPrimaryKey(&tt.user)
		assert.Equal(t, nil, err)
	}
}
