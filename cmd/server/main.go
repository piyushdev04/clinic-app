package main

import (
	_ "clinic-app/docs"
	"clinic-app/internal/auth"
	"clinic-app/internal/config"
	"clinic-app/internal/db"
	"clinic-app/internal/handlers"
	"clinic-app/internal/middleware"
	"clinic-app/internal/models"
	"clinic-app/internal/repositories"
	"clinic-app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load .env file
	_ = godotenv.Load()
	cfg := config.Load()
	database := db.Connect(cfg)
	dbAutoMigrate(database)
	db.SeedUsers(database)

	auth.SetJWTKey(cfg.JWTSecret)

	userRepo := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepo)
	patientRepo := repositories.NewPatientRepository(database)
	patientService := services.NewPatientService(patientRepo)
	patientHandler := handlers.NewPatientHandler(patientService)

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler(userService))

	authGroup := r.Group("/")
	authGroup.Use(middleware.JWTAuth())
	{
		authGroup.POST("/patients", middleware.Role("receptionist"), patientHandler.CreatePatient)
		authGroup.GET("/patients", patientHandler.GetPatients)
		authGroup.GET("/patients/:id", patientHandler.GetPatient)
		authGroup.PUT("/patients/:id", middleware.Role("receptionist"), patientHandler.UpdatePatient)
		authGroup.DELETE("/patients/:id", middleware.Role("receptionist"), patientHandler.DeletePatient)
		authGroup.PUT("/patients/:id/notes", middleware.Role("doctor"), patientHandler.UpdateNotes)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.ServerAddr)
}

func dbAutoMigrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{}, &models.Patient{})
}
