package web

import "time"

type BookResponse struct {
	BookID    int64     `json:"id"`
	BookName  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
