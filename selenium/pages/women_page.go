package pages

type WomenPage struct {
	Page Page
}

var (
	pantsAndDenium = ".catblocks>li:nth-of-type(3)>a"
)

func (s *WomenPage) SelectPantsAndDenim() *CategoryPage {
	s.Page.FindElementByCss(pantsAndDenium).Click()
	return &CategoryPage{Page:s.Page}
}


