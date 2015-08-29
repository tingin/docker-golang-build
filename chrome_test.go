package main

import (
	"github.com/tingin/webdriver"
	"log"
	"testing"
	"time"
)

func TestChrome(t *testing.T) {
	chromeDriver := webdriver.NewChromeDriver("/bin/chromedriver")
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
	time.Sleep(10 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
