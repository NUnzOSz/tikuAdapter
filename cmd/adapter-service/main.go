package main

import (
	"embed"
	"tikuAdapter/internal/api"
	"tikuAdapter/internal/registry/manager"
	"tikuAdapter/internal/service/timer"
	"tikuAdapter/pkg/logger"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
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

	go func() {
		server := gin.Default()
		api.SetAPIRouter(server)
		// Uncomment the following line if you need to serve web content
		api.SetWebRouter(buildFS, indexPage, server)
		timer.StartTimer()
		if err := server.Run("0.0.0.0:8060"); err != nil {
			logger.FatalLog(err)
		}
	}()

	systray.Run(func() {
		// Load your tray icon here (you need to provide the icon data)
		// var iconData []byte // This should be the raw data of your icon
		systray.SetIcon(icon.Data)
		systray.SetTemplateIcon(icon.Data, icon.Data)

		// For simplicity, we'll omit loading the icon in this example
		quit := systray.AddMenuItem("Quit", "Quit the application")
		go func() {
			<-quit.ClickedCh
			systray.Quit()
		}()
	}, func() {
		// Cleanup code (if any)
	})
}
