package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	
	"toba-lapor-backend/database"
	"toba-lapor-backend/internal/config"
	"toba-lapor-backend/internal/handler"
	"toba-lapor-backend/internal/middleware"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/internal/service"
	"toba-lapor-backend/internal/usecase"
)

func main() {
	// 1. Connect ke Database
	config.ConnectDatabase()

	// 2. Run Migration & Seeder
	database.RunMigration()
	database.RunSeeder()

	// 2.5 Connect to Firebase
	config.ConnectFirebase()

	// 3. Setup Dependencies
	db := config.DB

	// Repositories
	roleRepo := repository.NewRoleRepository(db)
	agencyRepo := repository.NewAgencyRepository(db)
	userRepo := repository.NewUserRepository(db)
	reportRepo := repository.NewReportRepository(db)
	historyRepo := repository.NewReportHistoryRepository(db)
	notifRepo := repository.NewNotificationRepository(db)

	// Services
	firebaseSvc := service.NewFirebaseService()

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, roleRepo)
	agencyUsecase := usecase.NewAgencyUsecase(agencyRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, roleRepo, agencyRepo)
	reportUsecase := usecase.NewReportUsecase(reportRepo, historyRepo, agencyRepo, notifRepo, userRepo, firebaseSvc)
	notifUsecase := usecase.NewNotificationUsecase(notifRepo)
	dashboardUsecase := usecase.NewDashboardUsecase(reportRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	agencyHandler := handler.NewAgencyHandler(agencyUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	reportHandler := handler.NewReportHandler(reportUsecase)
	dashboardHandler := handler.NewDashboardHandler(dashboardUsecase)
	notifHandler := handler.NewNotificationHandler(notifUsecase)

	// 4. Setup Gin Router
	r := gin.Default()

	// Serve static files from "uploads" directory
	r.Static("/uploads", "./uploads")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to TobaLapor API"})
	})

	// Public Routes
	api := r.Group("/api")
	api.POST("/auth/login", authHandler.Login)
	api.POST("/auth/register", authHandler.Register)

	// Protected Routes (User / Masyarakat)
	userRoutes := api.Group("")
	userRoutes.Use(middleware.AuthMiddleware("user"))
	{
		// Profile
		userRoutes.GET("/profile", userHandler.GetProfile)
		userRoutes.PUT("/profile", userHandler.UpdateProfile)
		userRoutes.PUT("/profile/fcm-token", userHandler.UpdateFCMToken)

		// Reports
		userRoutes.POST("/reports", reportHandler.Create)
		userRoutes.GET("/reports/me", reportHandler.GetMyReports)
		userRoutes.GET("/reports/:id", reportHandler.GetDetail)

		// Notifications
		userRoutes.GET("/notifications", notifHandler.GetMyNotifications)
		userRoutes.PUT("/notifications/:id/read", notifHandler.MarkAsRead)
	}

	// Protected Routes (Super Admin)
	superAdminRoutes := api.Group("/superadmin")
	superAdminRoutes.Use(middleware.AuthMiddleware("super_admin"))
	{
		// Agency Management
		superAdminRoutes.POST("/agencies", agencyHandler.Create)
		superAdminRoutes.GET("/agencies", agencyHandler.GetAll)
		superAdminRoutes.PUT("/agencies/:id", agencyHandler.Update)
		superAdminRoutes.DELETE("/agencies/:id", agencyHandler.Delete)

		// Admin Dinas Management
		superAdminRoutes.POST("/admins", userHandler.CreateAdminDinas)
		superAdminRoutes.GET("/admins", userHandler.GetAllAdminDinas)
		superAdminRoutes.PUT("/admins/:id", userHandler.UpdateAdminDinas)

		// User/Masyarakat Management
		superAdminRoutes.GET("/users", userHandler.GetAllUsers)
		superAdminRoutes.PATCH("/users/:id/status", userHandler.ToggleUserStatus)

		// Report Management
		superAdminRoutes.GET("/reports", reportHandler.GetAllReports)
		superAdminRoutes.PUT("/reports/:id/verify", reportHandler.VerifyReport)
		superAdminRoutes.PUT("/reports/:id/reject", reportHandler.RejectReport)

		// Dashboard
		superAdminRoutes.GET("/dashboard", dashboardHandler.GetStats)
	}

	// 5. Jalankan Server
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}