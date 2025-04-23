package routes

import (
        "github.com/gin-gonic/gin"
        "go.uber.org/zap"
        "gorm.io/gorm"

        "hospital-portal/internal/auth"
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

        // Initialize repositories
        userRepo := repositories.NewUserRepository(db)
        patientRepo := repositories.NewPatientRepository(db)

        // Initialize services
        authService := services.NewAuthService(userRepo, logger)
        patientService := services.NewPatientService(patientRepo, logger)

        // Initialize controllers
        authController := controllers.NewAuthController(authService, logger)
        patientController := controllers.NewPatientController(patientService, logger)

        // Auth routes
        r.POST("/api/login", authController.Login)
        r.POST("/api/register", authController.Register)

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
                        patients.PUT("/:id", middlewares.RoleMiddleware(auth.RoleDoctor), patientController.UpdatePatient)
                        
                        // Routes only available to receptionists
                        receptionistGroup := patients.Group("")
                        receptionistGroup.Use(middlewares.RoleMiddleware(auth.RoleReceptionist))
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
                        "endpoints": []string{
                                "/api/login - User login",
                                "/api/register - User registration",
                                "/api/v1/patients - Patient management (requires authentication)",
                                "/health - Server health check",
                        },
                })
        })
}
