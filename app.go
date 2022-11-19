package main

import (
	"context"
	_ "embed"
	"fmt"
	"io/ioutil"

	"github.com/KnockOutEZ/Kodee/backend/systemUsage"
	"github.com/KnockOutEZ/Kodee/backend/utils"
	"github.com/KnockOutEZ/Kodee/backend/auth"
	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var myCtx context.Context

//go:embed frontend/src/assets/wails.ico
var logoICO []byte

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	myCtx = ctx
	utils.CopyIconInStartup()
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(logoICO)
	systray.SetTitle("Kodee")
	systray.SetTooltip("Kodee-Your Personal Assistant")
	mOpen := systray.AddMenuItem("Open", "Open the app")
	go func() {
		for{
		<-mOpen.ClickedCh
		runtime.Show(myCtx)
		// runtime.WindowShow(myCtx)
		}
	}()
	mQuit := systray.AddMenuItem("Quit", "Quit the app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		runtime.Quit(myCtx)
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	utils.CheckErr(err)
	return b
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Meet returns a greeting for the given name
func (a *App) Notification(title,message string) {
	utils.NotificationFunc(title,message)
}


//cpu usage
func (a *App) GetCpuUsage() string{
	return systemUsage.GetCpuUsage()
}


//ram usage
func (a *App) GetRamUsage() []string{
	return systemUsage.GetRamUsage()
}

func (a *App) GetBandwithSpeed() []interface{}{
	return systemUsage.GetBandwithSpeed()
}

// Greet returns a greeting for the given name
func (a *App) Auth() {
	auth.HandlePage()
}