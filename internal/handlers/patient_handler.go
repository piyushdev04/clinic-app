package handlers

import (
	"clinic-app/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define an interface for the patient service
// Place this at the top of the file, or import from services if already defined

type PatientServiceInterface interface {
	CreatePatient(patient *models.Patient) error
	GetAllPatients() ([]models.Patient, error)
	GetPatientByID(id uint) (*models.Patient, error)
	UpdatePatient(patient *models.Patient) error
	DeletePatient(id uint) error
}

type PatientHandler struct {
	service PatientServiceInterface
}

func NewPatientHandler(service PatientServiceInterface) *PatientHandler {
	return &PatientHandler{service}
}

// @Security BearerAuth
// CreatePatient godoc
// @Summary Create Patient
// @Tags Receptionist
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Router /patients [post]
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := h.service.CreatePatient(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

// @Security BearerAuth
// GetPatients godoc
// @Summary List Patients
// @Tags Receptionist,Doctor
// @Produce json
// @Success 200 {array} models.Patient
// @Router /patients [get]
func (h *PatientHandler) GetPatients(c *gin.Context) {
	patients, err := h.service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

// @Security BearerAuth
// GetPatient godoc
// @Summary Get Patient
// @Tags Receptionist,Doctor
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Router /patients/{id} [get]
func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	patient, err := h.service.GetPatientByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

// @Security BearerAuth
// UpdatePatient godoc
// @Summary Update Patient
// @Tags Receptionist
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Patient data"
// @Success 200 {object} models.Patient
// @Router /patients/{id} [put]
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	patient.ID = uint(id)
	if err := h.service.UpdatePatient(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}

// @Security BearerAuth
// DeletePatient godoc
// @Summary Delete Patient
// @Tags Receptionist
// @Param id path int true "Patient ID"
// @Success 204
// @Router /patients/{id} [delete]
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeletePatient(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Security BearerAuth
// UpdateNotes godoc
// @Summary Update Patient Notes
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param notes body string true "Notes"
// @Success 200 {object} models.Patient
// @Router /patients/{id}/notes [put]
func (h *PatientHandler) UpdateNotes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Notes string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	patient, err := h.service.GetPatientByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	patient.Notes = req.Notes
	if err := h.service.UpdatePatient(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}
