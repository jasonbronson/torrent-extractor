package main

import (
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	site := os.Getenv("SITE")
	//remote := os.Getenv("REMOTE")
	// l := launcher.MustNewManaged(remote)
	// l.Set("disable-gpu").Delete("disable-gpu")
	// l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	// browser := rod.New().Client(l.MustClient()).MustConnect().MustPage(site)
	u := launcher.New().Logger(os.Stdout).Leakless(false).Bin("/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary")
	u.Flags["disable-gpu"] = nil
	b := u.MustLaunch()
	browser := rod.New().NoDefaultDevice().ControlURL(b).MustConnect().MustPage(site)
	browser.MustWindowFullscreen()
	prop := proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
	}
	browser.SetUserAgent(&prop)
	c := make([]*proto.NetworkCookieParam, 0)
	cookieparam := proto.NetworkCookieParam{
		Name:  "cf_clearance",
		Value: "Ep1OdpAycOaSv_R0XqMHOaA41ZQ93pyCiLaJ4IVV7qY-1652202056-0-150",
	}
	c = append(c, &cookieparam)
	browser.SetCookies(c)
	browser.MustElementX("/html/body/div[1]/img")
	browser.MustWaitLoad().MustScreenshot("a.png")
	//utils.Pause()
	log.Println("done")
	time.Sleep(time.Hour)
}
