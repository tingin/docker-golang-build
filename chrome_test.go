package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"testing"
	"time"
)

func TestChrome(t *testing.T) {
	log.Println("start")
	chromeDriver := webdriver.NewChromeDriver("chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("http://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(1 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
