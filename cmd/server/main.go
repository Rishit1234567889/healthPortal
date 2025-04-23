package main

import (
        "fmt"
        "log"
        "os"

        "github.com/gin-gonic/gin"
        "github.com/spf13/viper"
        "go.uber.org/zap"
        "gorm.io/driver/postgres"
        "gorm.io/gorm"

        "hospital-portal/internal/models"
        "hospital-portal/internal/routes"
)

func initConfig() {
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.AddConfigPath("./configs")
        viper.AutomaticEnv()

        if err := viper.ReadInConfig(); err != nil {
                log.Fatalf("Error reading config file: %s", err)
        }

        // Override with environment variables if they exist
        if os.Getenv("PORT") != "" {
                viper.Set("server.port", os.Getenv("PORT"))
        }

        if os.Getenv("JWT_SECRET") != "" {
                viper.Set("auth.jwt_secret", os.Getenv("JWT_SECRET"))
        }
}

func initDB() *gorm.DB {
        var dsn string

        // Check if DATABASE_URL is set (Replit PostgreSQL)
        if os.Getenv("DATABASE_URL") != "" {
                // Use the DATABASE_URL directly
                dsn = os.Getenv("DATABASE_URL")
        } else {
                // Construct DSN from individual environment variables or config
                dsn = fmt.Sprintf(
                        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
                        getEnvOrDefault("PGHOST", viper.GetString("database.host")),
                        getEnvOrDefault("PGUSER", viper.GetString("database.user")),
                        getEnvOrDefault("PGPASSWORD", viper.GetString("database.password")),
                        getEnvOrDefault("PGDATABASE", viper.GetString("database.dbname")),
                        getEnvOrDefault("PGPORT", viper.GetString("database.port")),
                )
        }

        log.Println("Connecting to database...")
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
                log.Fatalf("Failed to connect to database: %v", err)
        }

        // Auto migrate the schema
        log.Println("Running auto migrations...")
        err = db.AutoMigrate(&models.User{}, &models.Patient{})
        if err != nil {
                log.Fatalf("Failed to migrate database: %v", err)
        }

        return db
}

func getEnvOrDefault(env string, defaultValue string) string {
        if value := os.Getenv(env); value != "" {
                return value
        }
        return defaultValue
}

func initLogger() *zap.Logger {
        logger, err := zap.NewProduction()
        if err != nil {
                log.Fatalf("Can't initialize zap logger: %v", err)
        }
        return logger
}

func main() {
        // Initialize configuration
        initConfig()

        // Initialize logger
        logger := initLogger()
        defer logger.Sync()

        // Initialize database connection
        db := initDB()

        // Set up Gin
        gin.SetMode(viper.GetString("server.mode"))
        r := gin.Default()

        // Setup routes
        routes.SetupRoutes(r, db, logger)

        // Start server
        port := viper.GetString("server.port")
        if port == "" {
                port = "8000" // Default port
        }

        logger.Info(fmt.Sprintf("Starting server on port %s", port))
        if err := r.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
                logger.Fatal("Failed to start server", zap.Error(err))
        }
}
