package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/kaimu/speedtest/providers/netflix"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/toast.v1"
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
	var checkErr = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
		dirname, err := os.UserHomeDir()
		checkErr(err)
		in, err := os.Open(dirname+`\Desktop\kodee.lnk`)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(dirname+`\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\kodee.lnk`)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatal(err)
	}
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

//cpu usage
func (a *App) GetCpuUsage() string{
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return err.Error()
	}
	usedPercent := fmt.Sprintf("%.2f", cpuPercent[0])
	return usedPercent + "%"
}


//ram usage
func (a *App) GetRamUsage() []string{
	m, err := mem.VirtualMemory()
	if err != nil {
		return []string{err.Error()}
	}
	usedMessage := fmt.Sprintf(
		"%s (%.2f%%)",
		getReadableSize(m.Used),
		m.UsedPercent,
	)
	return []string{usedMessage, getReadableSize(m.Total),getReadableSize(m.Available),getReadableSize(m.Free)}
}
func getReadableSize(sizeInBytes uint64) (readableSizeString string) {
	var (
		units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units) && size >= 1024; i++ {
		size = size / 1024
	}
	readableSizeString = fmt.Sprintf("%.2f %s", size, units[i])
	return
}

func (a *App) GetBandwithSpeed() []interface{}{
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	netflixServer,_ := netflix.Fetch()
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

	return []interface{}{s.Latency, s.DLSpeed, s.ULSpeed,netflixServer}
	}

	return nil
}