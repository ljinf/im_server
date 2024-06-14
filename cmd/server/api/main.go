package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ljinf/im_server/cmd/server/api/wire"

	"github.com/ljinf/im_server/pkg/config"
	"github.com/ljinf/im_server/pkg/log"
	"go.uber.org/zap"
)

func main() {
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host",
		fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
