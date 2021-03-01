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

	hostPort := fmt.Sprintf("%s:%s", opt.Config.Server.Host, opt.Config.Server.Port)
	logger.Info(fn, zap.String("info", hostPort))

	if opt.Config.Server.Framework == "fiber" {
		r := routes.SetupFiber(routes.SetupOptFiber{
			Config: opt.Config,
		})

		err := r.Listen(hostPort)
		if err != nil {
			logger.Fatal(fn, zap.Error(err))
		}
	} else if opt.Config.Server.Framework == "gin" {
		r, err := routes.SetupGin(routes.SetupOptGin{
			Config: opt.Config,
		})
		if err != nil {
			logger.Fatal(fn, zap.Error(err))
		}
		err = r.Run(hostPort)
		if err != nil {
			logger.Fatal(fn, zap.Error(err))
		}
	} else if opt.Config.Server.Framework == "echo" {
		r := routes.SetupEcho(routes.SetupOptEcho{
			Config: opt.Config,
		})

		err := r.Start(hostPort)
		if err != nil {
			logger.Fatal(fn, zap.Error(err))
		}
	}
}
