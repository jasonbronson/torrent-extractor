package main

import (
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	site := os.Getenv("SITE")
	l := launcher.MustNewManaged("")
	l.Set("disable-gpu").Delete("disable-gpu")
	l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	browser := rod.New().Client(l.MustClient()).MustConnect().MustPage(site)
	browser.MustElementX("/html/body/div[1]/img")
	browser.MustWaitLoad().MustScreenshot("a.png")
	//utils.Pause()
	log.Println("done")
}
