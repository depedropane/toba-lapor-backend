package database

import (
	"fmt"
	"os"

	"toba-lapor-backend/internal/config"
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/pkg/utils"
)

func RunSeeder() {
	fmt.Println("Running Seeder...")
	db := config.DB

	// 1. Seed Roles
	roles := []string{"super_admin", "admin_dinas", "user"}
	for _, roleName := range roles {
		var role model.Role
		err := db.Where("name = ?", roleName).First(&role).Error
		if err != nil {
			db.Create(&model.Role{Name: roleName})
		}
	}

	// 2. Seed Super Admin
	var superAdminRole model.Role
	db.Where("name = ?", "super_admin").First(&superAdminRole)

	adminEmail := os.Getenv("SUPERADMIN_EMAIL")
	adminPassword := os.Getenv("SUPERADMIN_PASSWORD")

	if adminEmail == "" || adminPassword == "" {
		fmt.Println("Warning: SUPERADMIN_EMAIL or SUPERADMIN_PASSWORD is not set in .env")
		fmt.Println("Skipping Super Admin seeder...")
		return
	}

	var adminUser model.User
	err := db.Where("email = ?", adminEmail).First(&adminUser).Error
	if err != nil {
		hash, _ := utils.HashPassword(adminPassword)
		db.Create(&model.User{
			Name:     "Super Admin",
			Email:    adminEmail,
			Password: hash,
			RoleID:   superAdminRole.ID,
		})
		fmt.Println("Super admin created:", adminEmail)
	}
	
	fmt.Println("Seeder completed!")
}
