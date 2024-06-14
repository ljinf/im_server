package repository

import (
	"context"
	"github.com/ljinf/im_server/internal/model"
	"github.com/ljinf/im_server/internal/repository"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AccountRepository interface {
	CreateAccountInfo(info *model.AccountInfo) error
	GetAccountInfoById(userId int64) (*model.Register, error)
	GetAccountInfo(phone, email string) (*model.Register, error)
	UpdateAccountInfo(info *model.Register) error

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

func (a *accountRepository) CreateAccountInfo(info *model.AccountInfo) error {
	return a.DB(context.Background()).Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		//注册信息
		registerInfo := model.Register{
			UserId:    info.UserId,
			Phone:     info.Phone,
			Email:     info.Email,
			Password:  info.Password,
			Salt:      info.Salt,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := tx.Create(&registerInfo).Error; err != nil {
			return err
		}

		//创建新用户信息
		userInfo := model.UserInfo{
			UserId:    info.UserId,
			Status:    1,
			CreatedAt: now,
			UpdatedAt: now,
		}
		return tx.Create(&userInfo).Error
	})
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

func (a *accountRepository) GetUserInfo(userId int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := a.DB(context.Background()).Where("user_id=?", userId).First(&userInfo).Error
	return &userInfo, err
}

func (a *accountRepository) UpdateUserInfo(info *model.UserInfo) error {
	return a.DB(context.Background()).Where("user_id=?", info.UserId).Save(info).Error
}
