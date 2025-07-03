package services

import (
	"clinic-app/internal/models"
	"clinic-app/internal/repositories"
)

type PatientService struct {
	repo repositories.PatientRepository
}

func NewPatientService(repo repositories.PatientRepository) *PatientService {
	return &PatientService{repo}
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	return s.repo.Create(patient)
}

func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *PatientService) UpdatePatient(patient *models.Patient) error {
	return s.repo.Update(patient)
}

func (s *PatientService) DeletePatient(id uint) error {
	return s.repo.Delete(id)
}
