package model

import "time"

type Report struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	AgencyID    *uint     // Nullable saat awal dibuat
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Location    string    `gorm:"type:text;not null"`
	Latitude    float64   `gorm:"type:decimal(10,8)"`
	Longitude   float64   `gorm:"type:decimal(11,8)"`
	Status      string    `gorm:"type:varchar(50);default:'Menunggu Verifikasi'"` // Menunggu Verifikasi, Diteruskan ke Dinas, Sedang Diproses, Selesai, Ditolak
	CreatedAt   time.Time
	UpdatedAt   time.Time
	
	User            User            `gorm:"foreignKey:UserID"`
	Agency          *Agency         `gorm:"foreignKey:AgencyID"`
	ReportImages    []ReportImage   `gorm:"foreignKey:ReportID"`
	ReportHistories []ReportHistory `gorm:"foreignKey:ReportID"`
}
