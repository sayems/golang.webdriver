package main

import (
	"testing"
	"github.com/sayems/golang.examples/selenium/pages"
)

func TestSelenium(t *testing.T) {

	login := pages.HomePage{Page:page}
	login.GoToAccountPage().
		CreateAnAccount().
		Register()

}