package entity

type WebPageDataRsp struct {
	TemplateName string
	PageData     any
}

func (w *WebPageDataRsp) GetTemplateName() string {
	return w.TemplateName
}

func (w *WebPageDataRsp) GetPageData() any {
	return w.PageData
}
