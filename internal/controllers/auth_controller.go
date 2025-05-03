package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hospital-portal/internal/auth"
	"hospital-portal/internal/models"
	"hospital-portal/internal/services"
	"hospital-portal/internal/utils"
)

// AuthController handles authentication related requests
type AuthController struct {
	authService *services.AuthService
	logger      *zap.Logger
}

// NewAuthController creates a new auth controller instance
func NewAuthController(authService *services.AuthService, logger *zap.Logger) *AuthController {
	return &AuthController{
		authService: authService,
		logger:      logger,
	}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse represents the login response body
type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Login handles user login
func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid login request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := c.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.logger.Error("Authentication failed", zap.Error(err), zap.String("email", req.Email))
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	// Determine the role
	var role auth.Role
	switch user.Role {
	case "doctor":
		role = auth.RoleDoctor
	case "receptionist":
		role = auth.RoleReceptionist
	case "patient":
		role = auth.RolePatient
	case "admin":
		role = auth.RoleAdmin
	default:
		c.logger.Error("Invalid user role", zap.String("role", user.Role))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid user role", nil)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, role)
	if err != nil {
		c.logger.Error("Failed to generate token", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	// Mask sensitive information
	user.Password = ""

	ctx.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  *user,
	})
}

// LoginDoctor handles doctor login
func (c *AuthController) LoginDoctor(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid doctor login request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := c.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.logger.Error("Doctor authentication failed", zap.Error(err), zap.String("email", req.Email))
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	// Check if user is a doctor
	if user.Role != "doctor" {
		c.logger.Error("User is not a doctor", zap.String("email", req.Email), zap.String("role", user.Role))
		utils.ErrorResponse(ctx, http.StatusForbidden, "Access denied: User is not a doctor", nil)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, auth.RoleDoctor)
	if err != nil {
		c.logger.Error("Failed to generate token", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	// Mask sensitive information
	user.Password = ""

	ctx.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  *user,
	})
}

// LoginReceptionist handles receptionist login
func (c *AuthController) LoginReceptionist(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid receptionist login request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := c.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.logger.Error("Receptionist authentication failed", zap.Error(err), zap.String("email", req.Email))
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	// Check if user is a receptionist
	if user.Role != "receptionist" {
		c.logger.Error("User is not a receptionist", zap.String("email", req.Email), zap.String("role", user.Role))
		utils.ErrorResponse(ctx, http.StatusForbidden, "Access denied: User is not a receptionist", nil)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, auth.RoleReceptionist)
	if err != nil {
		c.logger.Error("Failed to generate token", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	// Mask sensitive information
	user.Password = ""

	ctx.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  *user,
	})
}

// LoginPatient handles patient login
func (c *AuthController) LoginPatient(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid patient login request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := c.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.logger.Error("Patient authentication failed", zap.Error(err), zap.String("email", req.Email))
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	// Check if user is a patient
	if user.Role != "patient" {
		c.logger.Error("User is not a patient", zap.String("email", req.Email), zap.String("role", user.Role))
		utils.ErrorResponse(ctx, http.StatusForbidden, "Access denied: User is not a patient", nil)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, auth.RolePatient)
	if err != nil {
		c.logger.Error("Failed to generate token", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	// Mask sensitive information
	user.Password = ""

	ctx.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  *user,
	})
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=doctor receptionist patient admin"`
}

// Register handles user registration
func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid registration request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create a new user
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Register the user
	createdUser, err := c.authService.RegisterUser(user, req.Password)
	if err != nil {
		c.logger.Error("Failed to register user", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	// Mask sensitive information
	createdUser.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    createdUser,
	})
}

// RegisterDoctor handles doctor registration
func (c *AuthController) RegisterDoctor(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid doctor registration request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Force the role to be doctor
	req.Role = "doctor"

	// Create a new user
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Register the user
	createdUser, err := c.authService.RegisterUser(user, req.Password)
	if err != nil {
		c.logger.Error("Failed to register doctor", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register doctor", err)
		return
	}

	// Mask sensitive information
	createdUser.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Doctor registered successfully",
		"user":    createdUser,
	})
}

// RegisterReceptionist handles receptionist registration
func (c *AuthController) RegisterReceptionist(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid receptionist registration request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Force the role to be receptionist
	req.Role = "receptionist"

	// Create a new user
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Register the user
	createdUser, err := c.authService.RegisterUser(user, req.Password)
	if err != nil {
		c.logger.Error("Failed to register receptionist", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register receptionist", err)
		return
	}

	// Mask sensitive information
	createdUser.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Receptionist registered successfully",
		"user":    createdUser,
	})
}

// RegisterPatient handles patient registration
func (c *AuthController) RegisterPatient(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid patient registration request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Force the role to be patient
	req.Role = "patient"

	// Create a new user
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Register the user
	createdUser, err := c.authService.RegisterUser(user, req.Password)
	if err != nil {
		c.logger.Error("Failed to register patient", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register patient", err)
		return
	}

	// Mask sensitive information
	createdUser.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Patient registered successfully",
		"user":    createdUser,
	})
}
