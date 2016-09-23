package pages

import (
	"time"
)

type HomePage struct {
	Page Page
}

var (
	account = "#header > div > div:nth-of-type(2) > a:nth-of-type(3) > span:nth-of-type(2)"
	myAccount = "#header-account > div > ul > li:nth-of-type(1) > a"
	womenTab = "#nav > ol > li:nth-of-type(1) > a"
	viewAllWomen = "#nav > ol > li:nth-of-type(1) > ul > li:nth-of-type(1) > aa"

)

func (s *HomePage) GoToAccountPage() *AccountPage {
	s.Page.FindElementByCss(account).Click()
	s.Page.FindElementByCss(myAccount).Click()
	return &AccountPage{Page:s.Page}
}

func (s *HomePage) GoToMenPage() *MenPage {
	return &MenPage{Page:s.Page}
}

func (s *HomePage) GoToWomenPage() *WomenPage {
	s.Page.MouseHoverToElement(womenTab)
	time.Sleep(time.Millisecond * 100)
	s.Page.FindElementByCss(viewAllWomen).Click()
	return &WomenPage{Page:s.Page}
}

func (s *HomePage) GoToAccessoriesPage() *AccessoriesPage {
	return &AccessoriesPage{Page:s.Page}
}

func (s *HomePage) GoToSalePage() *SalePage {
	return &SalePage{Page:s.Page}
}
