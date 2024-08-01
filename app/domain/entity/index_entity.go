package entity

type IndexPageData struct {
	User    *User
	Wisdom  string
	Content string
}

type User struct {
	Name string
}

type IndexData struct {
	Content string
}
