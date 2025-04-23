package services

import (
	"errors"

	"go.uber.org/zap"

	"hospital-portal/internal/models"
	"hospital-portal/internal/repositories"
	"hospital-portal/internal/utils"
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo *repositories.UserRepository
	logger   *zap.Logger
}

// NewAuthService creates a new auth service instance
func NewAuthService(userRepo *repositories.UserRepository, logger *zap.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		logger:   logger,
	}
}

// RegisterUser registers a new user
func (s *AuthService) RegisterUser(user *models.User, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, err
	}
	
	// Set the hashed password
	user.Password = hashedPassword
	
	// Create the user
	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err), zap.String("email", user.Email))
		return nil, err
	}
	
	return createdUser, nil
}

// AuthenticateUser authenticates a user
func (s *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	// Find the user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		s.logger.Error("User not found", zap.Error(err), zap.String("email", email))
		return nil, errors.New("invalid email or password")
	}
	
	// Verify the password
	if !utils.CheckPasswordHash(password, user.Password) {
		s.logger.Warn("Password mismatch", zap.String("email", email))
		return nil, errors.New("invalid email or password")
	}
	
	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
