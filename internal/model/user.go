package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	RoleID    uint      `gorm:"not null"`
	AgencyID  *uint     // Nullable untuk user dan super_admin
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(20)"`
	FCMToken  string    `gorm:"type:varchar(255)"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	Role      Role      `gorm:"foreignKey:RoleID"`
	Agency    *Agency   `gorm:"foreignKey:AgencyID"`
}
