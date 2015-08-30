package main

import (
	"github.com/tingin/webdriver"
	"log"
	"testing"
	"time"
)

func TestChrome(t *testing.T) {
	log.Println("start")
	chromeDriver := webdriver.NewChromeDriverRemote("http://127.0.0.1:9515")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("http://www.safeuway.com/")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(10 * time.Second)
	session.Delete()
	//chromeDriver.Stop()
	log.Println("end")
}
