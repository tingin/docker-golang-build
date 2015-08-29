package main

import (
	"github.com/fedesog/webdriver"
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

	chromeDriver.Stop()
}
