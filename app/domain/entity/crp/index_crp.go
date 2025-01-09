package crp

// PageDataIndexRsp PageDataIndex 首页数据（协议）
type PageDataIndexRsp struct {
	User    *User
	Wisdom  string
	Content string
}

// User 用户信息
type User struct {
	Name string
}
