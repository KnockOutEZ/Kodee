package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "KODEE",
		Width:            1024,
		Height:           728,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown: app.shutdown,
		OnBeforeClose: app.beforeClose,
		OnDomReady: app.domReady,
		HideWindowOnClose: true,
		Bind: []interface{}{
			app,
		},
	})
	
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}