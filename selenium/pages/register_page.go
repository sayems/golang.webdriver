package pages

import "math/rand"

type RegisterPage struct {
	Page Page
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	firstName = "firstname"
	lastName = "lastname"
	emailAddress = "email_address"
	password = "password"
	confirmPassword = "confirmation"
	registerButton = ".button"
)

func (s *RegisterPage) Register() *AccountPage {
	s.Page.FindElementById(firstName).SendKeys("Test")
	s.Page.FindElementById(lastName).SendKeys("Automation")
	s.Page.FindElementById(emailAddress).SendKeys(random(10) + "@automation.com")
	s.Page.FindElementById(password).SendKeys("Passw0rd!")
	s.Page.FindElementById(confirmPassword).SendKeys("Passw0rd!")
	s.Page.FindElementByCss(registerButton).Click()
	return &AccountPage{Page:s.Page}
}

func random(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}