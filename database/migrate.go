package database

import (
	"fmt"
	"log"

	"toba-lapor-backend/internal/config"
	"toba-lapor-backend/internal/model"
)

func RunMigration() {
	fmt.Println("Running Auto Migration...")
	
	err := config.DB.AutoMigrate(
		&model.Role{},
		&model.Agency{},
		&model.User{},
		&model.Report{},
		&model.ReportImage{},
		&model.ReportHistory{},
		&model.Notification{},
	)
	
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	
	fmt.Println("Migration completed successfully!")
}
