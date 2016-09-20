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
	s.Page.FindById(firstName).SendKeys("Test")
	s.Page.FindById(lastName).SendKeys("Automation")
	s.Page.FindById(emailAddress).SendKeys(random(10) + "@automation.com")
	s.Page.FindById(password).SendKeys("Passw0rd!")
	s.Page.FindById(confirmPassword).SendKeys("Passw0rd!")
	s.Page.FindByCss(registerButton).Click()
	return &AccountPage{Page:s.Page}
}

func random(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}