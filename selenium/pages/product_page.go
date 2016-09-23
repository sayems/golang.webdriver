package pages

type ProductPage struct {
	Page Page
}

func (s *ProductPage) addToCart() *CheckoutPage  {

	return &CheckoutPage{Page:s.Page}
}





