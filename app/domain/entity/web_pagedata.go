package entity

type WebPageData struct {
	TemplateName string
	PageData     any
}

func (w *WebPageData) GetTemplateName() string {
	return w.TemplateName
}

func (w *WebPageData) GetPageData() any {
	return w.PageData
}
