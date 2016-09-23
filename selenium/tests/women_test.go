package tests

import (
	"testing"
	"github.com/sayems/golang.webdriver/selenium/pages"
)

func TestWomenPage(t *testing.T) {

	login := pages.HomePage{Page:page}
	login.GoToWomenPage()

}