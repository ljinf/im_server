package server

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/ljinf/im_server/api/v1"
	"github.com/ljinf/im_server/internal/api/handler"
	"github.com/ljinf/im_server/internal/middleware"
	"github.com/ljinf/im_server/pkg/jwt"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/ljinf/im_server/pkg/server/http"
	"github.com/spf13/viper"
)

func NewAccountApiServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	acHandler handler.AccountApiHandler,
) *http.Server {

	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)
	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)
	s.GET("/", func(ctx *gin.Context) {
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using imServer account api!",
		})
	})

	v1 := s.Group("/v1")
	{

		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", acHandler.Register)
			noAuthRouter.POST("/login", acHandler.Login)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.GET("/account/info", acHandler.GetAccount)
			strictAuthRouter.PUT("/account/edit", acHandler.UpdateAccount)
			strictAuthRouter.GET("/account/user/info", acHandler.GetUserInfo)
			strictAuthRouter.PUT("/account/user/edit", acHandler.UpdateUserInfo)
		}
	}

	return nil
}
