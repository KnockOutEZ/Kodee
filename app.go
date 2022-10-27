package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/toast.v1"
)

var myCtx context.Context

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
	var checkErr = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
		dirname, err := os.UserHomeDir()
		cpCmd := exec.Command("cp", "-rf", dirname+`\Desktop\kodee.lnk`, dirname+`\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`)
		err = cpCmd.Run()
		checkErr(err)
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	myCtx = ctx
	systray.Run(onReady, onExit)
}
func onReady() {
	// systray.SetIcon(getIcon("./src/assets/icon.png"))
	systray.SetTitle("I'm alive!")
	systray.SetTooltip("Look at me, I'm a tooltip!")
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
	if err != nil {
		fmt.Print(err)
	}
	return b
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Meet returns a greeting for the given name
func (a *App) Notification(title,message string) {
	notification := toast.Notification{
		AppID:               "Kodee",
		Title:               title,
		Message:             message,
		// Icon:                "",
		// Actions:             []toast.Action{{"protocol", "I'm a button", "https://www.google.com/search?q=qwe"}, {"protocol", "Me too!", ""}},
	}
    err := notification.Push()
    if err != nil {
        log.Fatalln(err)
    }
}