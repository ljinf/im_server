//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	handler2 "github.com/ljinf/im_server/internal/api/handler"
	server2 "github.com/ljinf/im_server/internal/api/server"
	service2 "github.com/ljinf/im_server/internal/api/service"
	"github.com/ljinf/im_server/internal/repository"
	repository2 "github.com/ljinf/im_server/internal/rpc/repository"
	"github.com/ljinf/im_server/pkg/app"
	"github.com/ljinf/im_server/pkg/jwt"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/ljinf/im_server/pkg/server/http"
	"github.com/ljinf/im_server/pkg/sid"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository2.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service2.NewService,
	service2.NewUserService,
)

var handlerSet = wire.NewSet(
	handler2.NewHandler,
	handler2.NewUserHandler,
)

var serverSet = wire.NewSet(
	server2.NewHTTPServer,
	server2.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server2.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
