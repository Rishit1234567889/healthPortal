package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	// "hospital-portal/internal/auth"
	"hospital-portal/internal/controllers"
	"hospital-portal/internal/middlewares"
	"hospital-portal/internal/repositories"
	"hospital-portal/internal/services"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine, db *gorm.DB, logger *zap.Logger) {
	// Middlewares
	r.Use(middlewares.LoggerMiddleware(logger))
	r.Use(gin.Recovery())

	// Serve static files from public directory
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
	r.Static("/img", "./public/img")

	// Serve admin dashboard
	r.StaticFile("/admin", "./public/admin.html")

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	patientRepo := repositories.NewPatientRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, logger)
	patientService := services.NewPatientService(patientRepo, logger)

	// Initialize controllers
	authController := controllers.NewAuthController(authService, logger)
	patientController := controllers.NewPatientController(patientService, logger)

	// Auth routes - General
	r.POST("/api/login", authController.Login)
	r.POST("/api/register", authController.Register)

	// Auth routes - Role-specific
	auth := r.Group("/api/auth")
	{
		// Doctor routes
		doctor := auth.Group("/doctor")
		{
			doctor.POST("/login", authController.LoginDoctor)
			doctor.POST("/register", authController.RegisterDoctor)
		}

		// Receptionist routes
		receptionist := auth.Group("/receptionist")
		{
			receptionist.POST("/login", authController.LoginReceptionist)
			receptionist.POST("/register", authController.RegisterReceptionist)
		}

		// Patient routes
		patient := auth.Group("/patient")
		{
			patient.POST("/login", authController.LoginPatient)
			patient.POST("/register", authController.RegisterPatient)
		}
	}

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Add authentication middleware to protected routes
		v1.Use(middlewares.AuthMiddleware(logger))

		// Patient routes
		patients := v1.Group("/patients")
		{
			// Routes available to both doctors and receptionists
			patients.GET("", patientController.GetAllPatients)
			patients.GET("/:id", patientController.GetPatientByID)

			// Routes only available to doctors
			patients.PUT("/:id", middlewares.RoleMiddleware("doctor"), patientController.UpdatePatient)

			// Routes only available to receptionists
			receptionistGroup := patients.Group("")
			receptionistGroup.Use(middlewares.RoleMiddleware("receptionist"))
			{
				receptionistGroup.POST("", patientController.CreatePatient)
				receptionistGroup.DELETE("/:id", patientController.DeletePatient)
			}
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Welcome page
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Hospital Portal API",
			"version": "1.0.0",
			"ui": []string{
				"/admin - Admin Dashboard with Dark Mode",
			},
			"documentation": "Use the /admin endpoint to access the admin dashboard with dark mode toggle.",
			"docker":        "Run 'make docker-up' to start the application using Docker.",
			"endpoints": gin.H{
				"general": []string{
					"/api/login - General user login",
					"/api/register - General user registration",
					"/health - Server health check",
				},
				"doctor": []string{
					"/api/auth/doctor/login - Doctor login",
					"/api/auth/doctor/register - Doctor registration",
				},
				"receptionist": []string{
					"/api/auth/receptionist/login - Receptionist login",
					"/api/auth/receptionist/register - Receptionist registration",
				},
				"patient": []string{
					"/api/auth/patient/login - Patient login",
					"/api/auth/patient/register - Patient registration",
				},
				"protected": []string{
					"/api/v1/patients - Patient management (requires authentication)",
				},
			},
		})
	})
}
