package model

import "time"

type ReportHistory struct {
	ID            uint      `gorm:"primaryKey"`
	ReportID      uint      `gorm:"not null"`
	UserID        uint      `gorm:"not null"` // Admin/Super Admin yang update
	Status        string    `gorm:"type:varchar(50);not null"`
	Notes         string    `gorm:"type:text"`
	ProofImageURL string    `gorm:"type:varchar(255)"` // Opsional
	CreatedAt     time.Time
	
	Report        Report    `gorm:"foreignKey:ReportID"`
	User          User      `gorm:"foreignKey:UserID"`
}
