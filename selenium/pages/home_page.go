package pages

type HomePage struct {
	Page Page
}

var (
	account = "#header > div > div:nth-of-type(2) > a:nth-of-type(3) > span:nth-of-type(2)"
	myAccount = "#header-account > div > ul > li:nth-of-type(1) > a"
)

func (s *HomePage) GoToAccountPage() *AccountPage {
	s.Page.FindByCss(account).Click()
	s.Page.FindByCss(myAccount).Click()
	return &AccountPage{Page:s.Page}
}

func (s *HomePage) GoToMenPage() *MenPage {
	return &MenPage{Page:s.Page}
}

func (s *HomePage) GoToWomenPage() *WomenPage {
	return &WomenPage{Page:s.Page}
}

func (s *HomePage) GoToAccessoriesPage() *AccessoriesPage {
	return &AccessoriesPage{Page:s.Page}
}

func (s *HomePage) GoToSalePage() *SalePage {
	return &SalePage{Page:s.Page}
}
