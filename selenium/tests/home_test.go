package tests

import (
	"testing"
	"github.com/sayems/golang.webdriver/selenium/pages"
)

func TestSelenium(t *testing.T) {

	login := pages.HomePage{Page:page}
	login.GoToAccountPage().
		CreateAnAccount().
		Register()

}