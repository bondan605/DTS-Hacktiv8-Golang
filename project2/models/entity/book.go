package entity

import "time"

type Book struct {
	BookID    int64  `gorm:"primaryKey"`
	BookName  string `gorm:"not null; type:varchar(50)"`
	Author    string `gorm:"not null; type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
