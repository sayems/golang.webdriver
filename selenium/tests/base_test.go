package main

import (
	"testing"
	"flag"
	"os"
	"github.com/tebeka/selenium"
	"fmt"
	"github.com/sayems/golang.examples/selenium/pages"
)

var driver selenium.WebDriver
var page pages.Page

func TestMain(m *testing.M) {

	var err error
	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	// remote to selenium server
	if driver, err = selenium.NewRemote(caps, ""); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}

	page = pages.Page{Driver:driver}
	err = driver.Get("http://magento-demo.lexiconn.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	flag.Parse()
	r := m.Run()

	driver.Quit()

	os.Exit(r)
}
