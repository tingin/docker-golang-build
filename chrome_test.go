package main

import (
	"github.com/tingin/webdriver"
	"log"
	"testing"
)

func TestChrome(t *testing.T) {
	log.Println("start")
	chromeDriver := webdriver.NewChromeDriver("chromedriver")
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
	err = session.Url("http://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	session.Delete()
	chromeDriver.Stop()
}
