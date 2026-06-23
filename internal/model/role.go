package model

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(50);not null;unique"` // super_admin, admin_dinas, user
	CreatedAt time.Time
	UpdatedAt time.Time
	
	Users     []User    `gorm:"foreignKey:RoleID"`
}
