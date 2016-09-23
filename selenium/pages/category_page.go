package pages

import (
	"time"
	"math/rand"
)

var (
	listOfProduct = ".products-grid.products-grid--max-4-col.first.last.odd > li > a"
)

type CategoryPage struct {
	Page Page
}

func (s *CategoryPage) AddProduct() *ProductPage{
	productList := s.Page.FindElementsByCss(listOfProduct)
	productList[rand.Intn(len(productList))].Click()
	time.Sleep(time.Millisecond * 5000)
	return &ProductPage{Page:s.Page}
}