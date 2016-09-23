package pages

import "time"

type ProductPage struct {
	Page Page
}

var (
	color = "#attribute92"
	colorOption = "#attribute92 > option:nth-of-type(2)"
	size = "#attribute180"
	sizeOption = "#attribute180 > option:nth-of-type(2)"
	addToCartBtn = ".button.btn-cart"
)

func (s *ProductPage) AddProductToCart() *CheckoutPage  {

	s.Page.FindElementByCss(color).Click()
	s.Page.FindElementByCss(colorOption).Click()
	s.Page.FindElementByCss(size).Click()
	s.Page.FindElementByCss(sizeOption).Click()
	s.Page.FindElementByCss(addToCartBtn).Click()
	time.Sleep(time.Millisecond * 5000)
	return &CheckoutPage{Page:s.Page}
}





