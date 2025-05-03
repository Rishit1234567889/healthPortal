package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// Role type for user roles
type Role string

const (
	RoleDoctor       Role = "doctor"
	RoleReceptionist Role = "receptionist"
	RolePatient      Role = "patient"
	RoleAdmin        Role = "admin"
)

// Claims represents the JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   Role   `json:"role"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(userID uint, email string, role Role) (string, error) {
	// Get token expiry from config
	tokenExpiry := viper.GetDuration("auth.token_expiry")
	if tokenExpiry == 0 {
		tokenExpiry = 24 * time.Hour // Default to 24 hours
	}

	// Create claims with user information
	claims := &Claims{
		UserID: userID,
		Role:   role,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiry).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hospital-portal",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get the JWT secret from configuration
	jwtSecret := viper.GetString("auth.jwt_secret")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not configured")
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	// Get the JWT secret from configuration
	jwtSecret := viper.GetString("auth.jwt_secret")
	if jwtSecret == "" {
		return nil, errors.New("JWT secret is not configured")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token and extract claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
