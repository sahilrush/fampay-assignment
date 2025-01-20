package models

import (
	"time"
)

// Video represents a YouTube video stored in the database
// type Video struct {
// 	ID          uint      `gorm:"primaryKey;autoIncrement"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	PublishedAt time.Time `json:"publishedAt"`
// 	Thumbnails  string    `json:"thumbnails"`
// }

type Video struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Thumbnails  string    `json:"thumbnails"`
}

// type Video struct {
// 	gorm.Model            // This embeds ID, CreatedAt, UpdatedAt, and DeletedAt
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	PublishedAt time.Time `json:"publishedAt"`
// 	Thumbnails  string    `json:"thumbnails"`
// }
