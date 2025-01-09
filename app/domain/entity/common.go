package entity

// PageLimit 页条件
type PageLimit struct {
	Random   bool // 是否需要随机
	Page     int
	PageSize int
}
