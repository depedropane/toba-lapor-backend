package model

import "time"

type Agency struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	
	Users       []User    `gorm:"foreignKey:AgencyID"`
	Reports     []Report  `gorm:"foreignKey:AgencyID"`
}
