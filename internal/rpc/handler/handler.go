package handler

import (
	"github.com/ljinf/im_server/pkg/log"
	"github.com/ljinf/im_server/pkg/sid"
)

type Handler struct {
	sid    *sid.Sid
	logger *log.Logger
}

func NewHandler(sid *sid.Sid, logger *log.Logger) *Handler {
	return &Handler{
		sid:    sid,
		logger: logger,
	}
}
