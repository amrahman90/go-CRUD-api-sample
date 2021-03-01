package main

import (
	"github.com/amrahman90/go-CRUD-api-sample/internal/appconfig"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/server"
)

func main() {
	config := appconfig.Init()

	server.Setup(server.SetupOpt{
		Config: config,
	})
}
