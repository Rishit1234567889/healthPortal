package services

import (
	"go.uber.org/zap"

	"hospital-portal/internal/models"
	"hospital-portal/internal/repositories"
)

// PatientService handles patient business logic
type PatientService struct {
	patientRepo *repositories.PatientRepository
	logger      *zap.Logger
}

// NewPatientService creates a new patient service instance
func NewPatientService(patientRepo *repositories.PatientRepository, logger *zap.Logger) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
		logger:      logger,
	}
}

// CreatePatient creates a new patient
func (s *PatientService) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	return s.patientRepo.Create(patient)
}

// GetAllPatients retrieves all patients
func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
	return s.patientRepo.FindAll()
}

// GetPatientByID retrieves a patient by ID
func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
	return s.patientRepo.FindByID(id)
}

// UpdatePatient updates a patient
func (s *PatientService) UpdatePatient(patient *models.Patient) (*models.Patient, error) {
	return s.patientRepo.Update(patient)
}

// DeletePatient deletes a patient
func (s *PatientService) DeletePatient(id uint) error {
	return s.patientRepo.Delete(id)
}
