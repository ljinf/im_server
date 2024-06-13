package repository

import (
	"context"
	"github.com/ljinf/im_server/internal/model"
	"github.com/ljinf/im_server/internal/repository"
	"strings"
)

type AccountRepository interface {
	CreateAccount(info *model.Register) error
	GetAccountInfoById(userId int64) (*model.Register, error)
	GetAccountInfo(phone, email string) (*model.Register, error)
	UpdateAccountInfo(info *model.Register) error
	CreateUserInfo(user *model.UserInfo) error
	GetUserInfo(userId int64) (*model.UserInfo, error)
	UpdateUserInfo(info *model.UserInfo) error
}

type accountRepository struct {
	*repository.Repository
}

func NewAccountRepository(repo *repository.Repository) AccountRepository {
	return &accountRepository{
		Repository: repo,
	}
}

func (a *accountRepository) CreateAccount(info *model.Register) error {
	return a.DB(context.Background()).Create(info).Error
}

func (a *accountRepository) GetAccountInfo(phone, email string) (*model.Register, error) {

	var conds []string
	if phone != "" {
		conds = append(conds, "phone =? ")
	}
	if email != "" {
		conds = append(conds, "email =? ")
	}

	var accountInfo model.Register
	err := a.DB(context.Background()).Where(strings.Join(conds, " and ")).First(&accountInfo).Error

	return &accountInfo, err

}

func (a *accountRepository) GetAccountInfoById(userId int64) (*model.Register, error) {
	var accountInfo model.Register
	err := a.DB(context.Background()).Where("user_id=?", userId).First(&accountInfo).Error
	return &accountInfo, err
}

func (a *accountRepository) UpdateAccountInfo(info *model.Register) error {
	return a.DB(context.Background()).Save(info).Error
}

func (a *accountRepository) CreateUserInfo(user *model.UserInfo) error {
	return a.DB(context.Background()).Save(user).Error
}

func (a *accountRepository) GetUserInfo(userId int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := a.DB(context.Background()).Where("user_id=?", userId).First(&userInfo).Error
	return &userInfo, err
}

func (a *accountRepository) UpdateUserInfo(info *model.UserInfo) error {
	return a.DB(context.Background()).Where("user_id=?", info.UserId).Save(info).Error
}
