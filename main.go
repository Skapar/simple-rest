package main

import (
	"github.com/Skapar/simple-rest/app"
)

func main() {
	// Create a new App
	app := app.NewApp()
	// defer app.DB.Close()

	app.Log.Infof("Starting server on port %s", app.Cfg.ListenHttpPort)
	// app.Run()

	return
}
