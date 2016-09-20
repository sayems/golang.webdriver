package pages

type AccountPage struct {
	Page Page
}

var (
	createAnAccount = "#login-form > div > div:nth-of-type(1) > div:nth-of-type(2) > a"
)

func (s *AccountPage) CreateAnAccount() *RegisterPage {
	s.Page.FindByCss(createAnAccount).Click()
	return &RegisterPage{Page:s.Page}
}