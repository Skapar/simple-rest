package main

import (
	"github.com/Skapar/simple-rest/app"
	_ "github.com/Skapar/simple-rest/docs"
)

func main() {
	// Create a new App
	application := app.NewApp()

	application.Log.Infof("Starting server on port %s", application.Cfg.ListenHttpPort)
	application.Log.Infof("Swagger UI available at http://localhost:%s/swagger/index.html", application.Cfg.ListenHttpPort)
	application.Log.Fatal(application.Router.Run(":" + application.Cfg.ListenHttpPort))
}
