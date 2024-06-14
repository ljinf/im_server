package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/ljinf/im_server/api/v1"
	"github.com/ljinf/im_server/internal/model"
	"github.com/ljinf/im_server/internal/rpc/repository"
	"github.com/ljinf/im_server/pkg/proto/account"
	pwd_encoder "github.com/ljinf/im_server/pkg/pwd-encoder"
	"go.uber.org/zap"
)

type AccountServerHandler struct {
	*Handler
	repo repository.AccountRepository
}

func NewAccountServerHandler(server *Handler, repo repository.AccountRepository) *AccountServerHandler {
	return &AccountServerHandler{
		Handler: server,
		repo:    repo,
	}
}

func (h *AccountServerHandler) CreateAccount(ctx context.Context, req *account.CreateAccountReq) (*account.CreateAccountRes, error) {

	userId, err := h.sid.GenUint64()
	if err != nil {
		h.logger.Error(err.Error())
		return nil, v1.ErrCreateIdFailed
	}
	if userId == 0 {
		h.logger.Error("err:create account return userId=0", zap.Any("userId", userId))
		return nil, v1.ErrCreateUserFailed
	}

	salt, encodePwd := pwd_encoder.PwdEncode(req.GetPassword())
	accountInfo := model.AccountInfo{
		UserId:   int64(userId),
		Phone:    req.Phone,
		Email:    req.Email,
		Password: encodePwd,
		Salt:     salt,
		Status:   1,
	}

	if err = h.repo.CreateAccountInfo(&accountInfo); err != nil {
		h.logger.Error(fmt.Sprintf("初始化用户信息失败 %v", err.Error()), zap.Any("user", accountInfo))
		return nil, v1.ErrCreateUserInfoFailed
	}

	return &account.CreateAccountRes{
		UserId: accountInfo.UserId,
		Phone:  accountInfo.Phone,
		Email:  accountInfo.Email,
	}, nil

}

func (h *AccountServerHandler) GetAccountInfo(ctx context.Context, req *account.AccountInfoReq) (*account.AccountInfoRes, error) {
	accountInfo, err := h.repo.GetAccountInfo(req.GetPhone(), req.GetEmail())
	if err != nil {
		h.logger.Error(err.Error(), zap.String("phone", req.Phone), zap.String("email", req.Phone))
		return nil, v1.ErrInternalServerError
	}

	return &account.AccountInfoRes{
		UserId:   accountInfo.UserId,
		Phone:    accountInfo.Phone,
		Email:    accountInfo.Email,
		Password: accountInfo.Password,
		Salt:     accountInfo.Salt,
	}, nil
}

func (h *AccountServerHandler) UpdateAccountInfo(ctx context.Context, req *account.UpdateAccountInfoReq) (*empty.Empty, error) {

	infoById, err := h.repo.GetAccountInfoById(req.UserId)
	if err != nil {
		h.logger.Error(err.Error(), zap.Any("userId", req.UserId))
		return nil, v1.ErrInternalServerError
	}

	//检查新的phone，email是否已被绑定
	if req.Phone != "" && req.Phone != infoById.Phone {
		infoByPhone, err := h.repo.GetAccountInfo(req.Phone, "")
		if err != nil {
			h.logger.Error(err.Error(), zap.Any("phone", req.Phone))
			return nil, v1.ErrInternalServerError
		}

		if infoByPhone != nil {
			return nil, v1.ErrPhoneAlreadyUse
		}

	}

	if req.Email != "" && req.Email != infoById.Email {
		infoByEmail, err := h.repo.GetAccountInfo("", req.Email)
		if err != nil {
			h.logger.Error(err.Error(), zap.Any("email", req.Email))
			return nil, v1.ErrInternalServerError
		}

		if infoByEmail != nil {
			return nil, v1.ErrEmailAlreadyUse
		}
	}

	//正常修改信息
	accountInfo := model.Register{
		UserId: req.UserId,
		Phone:  req.Phone,
		Email:  req.Email,
	}

	if req.Password != "" {
		salt, encodePwd := pwd_encoder.PwdEncode(req.GetPassword())
		accountInfo.Password = encodePwd
		accountInfo.Salt = salt
	}

	if err := h.repo.UpdateAccountInfo(&accountInfo); err != nil {
		h.logger.Error(err.Error(), zap.Any("accountInfo", accountInfo))
		return nil, v1.ErrUpdateUserInfoFailed
	}

	return &empty.Empty{}, nil

}

func (h *AccountServerHandler) GetUserInfo(ctx context.Context, req *account.UserInfoReq) (*account.UserInfoRes, error) {
	userInfo, err := h.repo.GetUserInfo(req.UserId)
	if err != nil {
		h.logger.Error(err.Error(), zap.Any("userInfo", userInfo))
		return nil, v1.ErrInternalServerError
	}

	res := &account.UserInfoRes{
		UserId:   userInfo.UserId,
		NickName: userInfo.NickName,
		Avatar:   userInfo.Avatar,
		Gender:   int32(userInfo.Gender),
	}

	return res, nil

}

func (h *AccountServerHandler) UpdateUserInfo(ctx context.Context, req *account.UpdateUserInfoReq) (*account.UpdateUserInfoRes, error) {
	info := model.UserInfo{
		UserId:   req.UserId,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}
	if err := h.repo.UpdateUserInfo(&info); err != nil {
		h.logger.Error(err.Error(), zap.Any("userInfo", info))
		return nil, v1.ErrUpdateUserInfoFailed
	}

	return &account.UpdateUserInfoRes{
		UserId:   req.UserId,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}, nil
}
