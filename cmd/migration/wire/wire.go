//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ljinf/im_server/internal/api/server"
	"github.com/ljinf/im_server/internal/repository"
	repository2 "github.com/ljinf/im_server/internal/rpc/repository"
	"github.com/ljinf/im_server/pkg/app"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository2.NewUserRepository,
)
var serverSet = wire.NewSet(
	server.NewMigrate,
)

// build App
func newApp(
	migrate *server.Migrate,
) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serverSet,
		newApp,
	))
}
