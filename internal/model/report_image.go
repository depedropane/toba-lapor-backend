package model

import "time"

type ReportImage struct {
	ID        uint      `gorm:"primaryKey"`
	ReportID  uint      `gorm:"not null"`
	ImageURL  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	
	Report    Report    `gorm:"foreignKey:ReportID"`
}
