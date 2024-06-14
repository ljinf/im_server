//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ljinf/im_server/internal/api/handler"
	"github.com/ljinf/im_server/internal/api/server"
	"github.com/ljinf/im_server/internal/api/service"
	"github.com/ljinf/im_server/pkg/app"
	"github.com/ljinf/im_server/pkg/jwt"
	"github.com/ljinf/im_server/pkg/log"
	"github.com/ljinf/im_server/pkg/server/http"
	"github.com/ljinf/im_server/pkg/sid"
	"github.com/spf13/viper"
)

var sidSet = wire.NewSet(
	sid.NewSid,
)

var jwtSet = wire.NewSet(
	jwt.NewJwt,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewAccountApiService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAccountApiHandler,
)

var serverSet = wire.NewSet(
	server.NewAccountApiServer,
)

// build App
func newApp(
	accountApi *http.Server,
) *app.App {
	return app.NewApp(
		app.WithServer(accountApi),
		app.WithName("api-servers"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		sidSet,
		jwtSet,
		serviceSet,
		handlerSet,
		serverSet,
		newApp,
	))
}
