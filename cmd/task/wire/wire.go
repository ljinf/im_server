//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ljinf/im_server/internal/api/server"
	"github.com/ljinf/im_server/pkg/app"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/spf13/viper"
)

var serverSet = wire.NewSet(
	server.NewTask,
)

// build App
func newApp(
	task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		newApp,
	))
}
