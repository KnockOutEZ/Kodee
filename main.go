package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
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
		Title:            "Kodee-Your Personal Assistant",
		Width:            1024,
		Height:           728,
		Assets:           assets,
		// BackgroundColour: &options.RGBA{R: 100, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown: app.shutdown,
		OnBeforeClose: app.beforeClose,
		OnDomReady: app.domReady,
		HideWindowOnClose: true,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
            WebviewIsTransparent:              true,
            WindowIsTranslucent:               false,
            Theme: windows.Light,
	}})
	
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}