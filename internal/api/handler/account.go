package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/im_server/internal/api/service"
)

type AccountApiHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
	UpdateAccount(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	UpdateUserInfo(ctx *gin.Context)
}

type accountApiHandler struct {
	*Handler
	srv service.AccountApiService
}

func NewAccountApiHandler(h *Handler, srv service.AccountApiService) AccountApiHandler {
	return &accountApiHandler{
		Handler: h,
		srv:     srv,
	}
}

func (h *accountApiHandler) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *accountApiHandler) Login(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *accountApiHandler) GetAccount(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *accountApiHandler) UpdateAccount(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *accountApiHandler) GetUserInfo(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *accountApiHandler) UpdateUserInfo(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
