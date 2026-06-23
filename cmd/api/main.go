package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/database"
	"toba-lapor-backend/internal/config"
	"toba-lapor-backend/internal/handler"
	"toba-lapor-backend/internal/middleware"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/internal/usecase"
)

func main() {
	// 1. Connect ke Database
	config.ConnectDatabase()

	// 2. Run Migration & Seeder
	database.RunMigration()
	database.RunSeeder()

	// 3. Setup Dependencies
	db := config.DB

	// Repositories
	roleRepo := repository.NewRoleRepository(db)
	agencyRepo := repository.NewAgencyRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo)
	agencyUsecase := usecase.NewAgencyUsecase(agencyRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, roleRepo, agencyRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	agencyHandler := handler.NewAgencyHandler(agencyUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	// 4. Setup Gin Router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to TobaLapor API"})
	})

	// Public Routes
	api := r.Group("/api")
	api.POST("/auth/login", authHandler.Login)

	// Protected Routes (Super Admin)
	superAdminRoutes := api.Group("/superadmin")
	superAdminRoutes.Use(middleware.AuthMiddleware("super_admin"))
	{
		// Agency Management
		superAdminRoutes.POST("/agencies", agencyHandler.Create)
		superAdminRoutes.GET("/agencies", agencyHandler.GetAll)
		superAdminRoutes.PUT("/agencies/:id", agencyHandler.Update)
		superAdminRoutes.DELETE("/agencies/:id", agencyHandler.Delete)

		superAdminRoutes.POST("/admins", userHandler.CreateAdminDinas)
		superAdminRoutes.GET("/admins", userHandler.GetAllAdminDinas)
		superAdminRoutes.PUT("/admins/:id", userHandler.UpdateAdminDinas)

		// User/Masyarakat Management
		superAdminRoutes.GET("/users", userHandler.GetAllUsers)
		superAdminRoutes.PATCH("/users/:id/status", userHandler.ToggleUserStatus)
	}

	// 5. Jalankan Server
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}