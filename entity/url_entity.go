package entity

import "time"

type URL struct {
	ID         uint   `gorm:"primaryKey"`
	Original   string `gorm:"not null"`
	ShortCode  string `gorm:"uniqueIndex;not null"`
	CreatedAt  time.Time
	ClickCount uint
}
