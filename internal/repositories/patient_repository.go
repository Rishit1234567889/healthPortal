package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"

	"hospital-portal/internal/models"
)

// PatientRepository handles database operations for patients
type PatientRepository struct {
	db *gorm.DB
}

// NewPatientRepository creates a new patient repository instance
func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{
		db: db,
	}
}

// Create creates a new patient
func (r *PatientRepository) Create(patient *models.Patient) (*models.Patient, error) {
	if err := r.db.Create(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

// FindAll retrieves all patients
func (r *PatientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	if err := r.db.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

// FindByID retrieves a patient by ID
func (r *PatientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.First(&patient, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, err
	}
	return &patient, nil
}

// FindByName retrieve a patient by name
// if needed
func (r *PatientRepository) FindByName(name string) (*models.Patient, error) {
	var patient models.Patient

	// Debugging log
	fmt.Println("Searching for patient in DB:", name)

	// Execute query
	err := r.db.Where("LOWER(name) = ?", strings.ToLower(name)).First(&patient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Patient not found in DB:", name) // Additional debugging log
			return nil, errors.New("patient not found")
		}

		return nil, err
	}

	return &patient, nil
}

// Update updates a patient
func (r *PatientRepository) Update(patient *models.Patient) (*models.Patient, error) {
	// Check if patient exists
	if err := r.db.First(&models.Patient{}, patient.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, err
	}

	// Update patient
	if err := r.db.Save(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

// Delete deletes a patient
func (r *PatientRepository) Delete(id uint) error {
	// Check if patient exists
	if err := r.db.First(&models.Patient{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("patient not found")
		}
		return err
	}

	// Delete patient
	if err := r.db.Delete(&models.Patient{}, id).Error; err != nil {
		return err
	}
	return nil
}
