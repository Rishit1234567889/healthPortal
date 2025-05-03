package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"

	"hospital-portal/internal/models"
	"hospital-portal/internal/services"
	"hospital-portal/internal/utils"
)

// PatientController handles patient-related requests
type PatientController struct {
	patientService *services.PatientService
	logger         *zap.Logger
}

// NewPatientController creates a new patient controller instance
func NewPatientController(patientService *services.PatientService, logger *zap.Logger) *PatientController {
	return &PatientController{
		patientService: patientService,
		logger:         logger,
	}
}

// PatientRequest represents the patient request body
type PatientRequest struct {
	Name           string `json:"name" binding:"required"`
	Age            int    `json:"age" binding:"required,min=0,max=150"`
	Gender         string `json:"gender" binding:"required,oneof=male female other"`
	Address        string `json:"address" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	MedicalHistory string `json:"medical_history"`
	Diagnosis      string `json:"diagnosis"`
	Treatment      string `json:"treatment"`
	Notes          string `json:"notes"`
}

// CreatePatient handles creating a new patient
func (c *PatientController) CreatePatient(ctx *gin.Context) {
	var req PatientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid patient create request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	patient := &models.Patient{
		Name:           req.Name,
		Age:            req.Age,
		Gender:         req.Gender,
		Address:        req.Address,
		PhoneNumber:    req.PhoneNumber,
		MedicalHistory: req.MedicalHistory,
		Diagnosis:      req.Diagnosis,
		Treatment:      req.Treatment,
		Notes:          req.Notes,
	}

	createdPatient, err := c.patientService.CreatePatient(patient)
	if err != nil {
		c.logger.Error("Failed to create patient", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create patient", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Patient created successfully",
		"patient": createdPatient,
	})
}

// GetAllPatients handles retrieving all patients
func (c *PatientController) GetAllPatients(ctx *gin.Context) {
	patients, err := c.patientService.GetAllPatients()
	if err != nil {
		c.logger.Error("Failed to fetch patients", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch patients", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"patients": patients,
	})
}

// GetPatientByID handles retrieving a patient by ID
func (c *PatientController) GetPatientByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.logger.Error("Invalid patient ID", zap.Error(err), zap.String("id", idStr))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	patient, err := c.patientService.GetPatientByID(uint(id))
	if err != nil {
		c.logger.Error("Failed to fetch patient", zap.Error(err), zap.Uint64("id", id))
		utils.ErrorResponse(ctx, http.StatusNotFound, "Patient not found", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"patient": patient,
	})
}

// GetPatientByName handles retrieving a patient by name
func (c *PatientController) GetPatientByName(ctx *gin.Context) {
	name := strings.TrimSpace(ctx.Param("name"))
	fmt.Println("Searching for:", name)

	if name == "" {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Name parameter is required", nil)
		return
	}

	patient, err := c.patientService.GetPatientByName(name)
	if err != nil {
		c.logger.Error("Failed to fetch patient", zap.Error(err), zap.String("name", name))
		utils.ErrorResponse(ctx, http.StatusNotFound, "Patient not found", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"patient": patient})
}

// UpdatePatient handles updating an existing patient
func (c *PatientController) UpdatePatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.logger.Error("Invalid patient ID", zap.Error(err), zap.String("id", idStr))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	var req PatientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("Invalid patient update request", zap.Error(err))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	patient := &models.Patient{
		Name:           req.Name,
		Age:            req.Age,
		Gender:         req.Gender,
		Address:        req.Address,
		PhoneNumber:    req.PhoneNumber,
		MedicalHistory: req.MedicalHistory,
		Diagnosis:      req.Diagnosis,
		Treatment:      req.Treatment,
		Notes:          req.Notes,
	}
	patient.ID = uint(id)

	updatedPatient, err := c.patientService.UpdatePatient(patient)
	if err != nil {
		c.logger.Error("Failed to update patient", zap.Error(err), zap.Uint64("id", id))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update patient", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Patient updated successfully",
		"patient": updatedPatient,
	})
}

// DeletePatient handles deleting a patient
func (c *PatientController) DeletePatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.logger.Error("Invalid patient ID", zap.Error(err), zap.String("id", idStr))
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	err = c.patientService.DeletePatient(uint(id))
	if err != nil {
		c.logger.Error("Failed to delete patient", zap.Error(err), zap.Uint64("id", id))
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete patient", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Patient deleted successfully",
	})
}
