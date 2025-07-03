package repositories

import (
	"clinic-app/internal/models"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetAll() ([]models.Patient, error)
	GetByID(id uint) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
}

type GormPatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &GormPatientRepository{db}
}

func (r *GormPatientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *GormPatientRepository) GetAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Find(&patients).Error
	return patients, err
}

func (r *GormPatientRepository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *GormPatientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *GormPatientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}
