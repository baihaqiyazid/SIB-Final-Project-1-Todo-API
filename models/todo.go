package models

import "time"

type Todo struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Text      string    `gorm:"not null;type:varchar" json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
