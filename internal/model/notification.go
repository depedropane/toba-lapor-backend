package model

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Body      string    `gorm:"type:text;not null"`
	IsRead    bool      `gorm:"default:false"`
	CreatedAt time.Time
	
	User      User      `gorm:"foreignKey:UserID"`
}
