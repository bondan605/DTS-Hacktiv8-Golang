package web

type BookRequest struct {
	BookName string `json:"name_book"`
	Author   string `json:"author"`
}
