package models

import (
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the hospital
type Patient struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"not null"`
	Age            int            `json:"age" gorm:"not null"`
	Gender         string         `json:"gender" gorm:"not null"`
	Address        string         `json:"address" gorm:"not null"`
	PhoneNumber    string         `json:"phone_number" gorm:"not null"`
	MedicalHistory string         `json:"medical_history"`
	Diagnosis      string         `json:"diagnosis"`
	Treatment      string         `json:"treatment"`
	Notes          string         `json:"notes"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
