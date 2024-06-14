package service

import "github.com/ljinf/im_server/internal/model"

type AccountApiService interface {
	CreateAccount(req *model.AccountInfo) error
	GetAccount(req *model.AccountInfo) (*model.AccountInfo, error)
	UpdateAccount(req *model.AccountInfo) (*model.AccountInfo, error)
	GetUserInfo(userId int64) (*model.UserInfo, error)
	UpdateUserInfo(req *model.AccountInfo) (*model.UserInfo, error)
}

type accountApiService struct {
	*Service
}

func NewAccountApiService(s *Service) AccountApiService {
	return &accountApiService{
		Service: s,
	}
}

func (s *accountApiService) CreateAccount(req *model.AccountInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *accountApiService) GetAccount(req *model.AccountInfo) (*model.AccountInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *accountApiService) UpdateAccount(req *model.AccountInfo) (*model.AccountInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *accountApiService) GetUserInfo(userId int64) (*model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *accountApiService) UpdateUserInfo(req *model.AccountInfo) (*model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}
