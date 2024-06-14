package service

import (
	"github.com/ljinf/im_server/pkg/jwt"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/ljinf/im_server/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
}

func NewService(
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
