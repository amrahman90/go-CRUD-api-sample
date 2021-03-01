package server

import (
	"fmt"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/routes"
	"go.uber.org/zap"
)

type SetupOpt struct {
	Config *config.Config
}

func Setup(opt SetupOpt) {
	fn := "server.Setup"
	logger := opt.Config.GetLogger()

	r := routes.Setup(routes.SetupOpt{
		Config: opt.Config,
	})

	hostPort := fmt.Sprintf("%s:%s", opt.Config.Server.Host, opt.Config.Server.Port)
	logger.Info(fn, zap.String("info", hostPort))
	err := r.Listen(hostPort)
	if err != nil {
		logger.Fatal(fn, zap.Error(err))
	}
}
