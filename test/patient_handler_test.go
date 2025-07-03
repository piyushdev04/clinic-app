package test

import (
	"bytes"
	"clinic-app/internal/handlers"
	"clinic-app/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockPatientService struct{}

func (m *mockPatientService) CreatePatient(patient *models.Patient) error     { return nil }
func (m *mockPatientService) GetAllPatients() ([]models.Patient, error)       { return nil, nil }
func (m *mockPatientService) GetPatientByID(id uint) (*models.Patient, error) { return nil, nil }
func (m *mockPatientService) UpdatePatient(patient *models.Patient) error     { return nil }
func (m *mockPatientService) DeletePatient(id uint) error                     { return nil }

func TestCreatePatient(t *testing.T) {
	svc := &mockPatientService{}
	h := handlers.NewPatientHandler(svc)

	patient := models.Patient{Name: "John", Age: 30}
	body, _ := json.Marshal(patient)
	req := httptest.NewRequest(http.MethodPost, "/patients", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	h.CreatePatient(c)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
