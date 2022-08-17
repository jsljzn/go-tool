package reqx

const DefaultPageIndex int = 1
const DefaultPageSize int = 50

type Page struct {
	PageIndex int `form:"page_index"`
	PageSize  int `form:"page_size"`
}

func (page *Page) GetPageIndex() int {
	if page.PageIndex < DefaultPageIndex {
		return (DefaultPageIndex - 1) * page.GetPageSize()
	}
	return (page.PageIndex - 1) * page.GetPageSize()
}

func (page *Page) GetPageSize() int {
	if page.PageSize >= DefaultPageSize {
		return DefaultPageSize
	}
	return page.PageSize
}
