package main

import (
	"embed"
	"tikuAdapter/internal/api"
	"tikuAdapter/internal/registry/manager"
	"tikuAdapter/internal/service/timer"
	"tikuAdapter/pkg/logger"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var buildFS embed.FS

//go:embed dist/index.html
var indexPage []byte

func main() {
	mg, err := manager.CreateManager()
	if err != nil {
		logger.FatalLog(err)
	}
	defer func(mg *manager.Manager) {
		err := mg.CloseManager()
		if err != nil {
			logger.FatalLog(err)
		}
	}(mg)

	server := gin.Default()
	api.SetAPIRouter(server)
	api.SetWebRouter(buildFS, indexPage, server)
	timer.StartTimer()
	err = server.Run("0.0.0.0:8060")
	if err != nil {
		logger.FatalLog(err)
	}
}
