package patients

import (
	"database/sql"

	patients "github.com/org-struct-hexagonal-architecture-/patients/models_domain"
)

// PatientGateway interface
type PatientGateway interface {
	CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error)
	GetPatients()
	GetPatientByID(id int64)
}

// CreatePatientInDB struct
type CreatePatientInDB struct {
	PatientStorage
}

// CreatePatient func
func (c *CreatePatientInDB) CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error) {
	return c.createPatientDB(p)
}

// GetPatients func
func (c *CreatePatientInDB) GetPatients() []*patients.Patient {
	return c.GetPatients()
}

// getPatientByID func
func (c *CreatePatientInDB) getPatientByID(id int64) (*patients.Patient, error) {
	return c.getPatientByIDDB(id)
}

// NewPatientGateway func
func NewPatientGateway(db *sql.DB) PatientGateway {
	return &CreatePatientInDB{NewPatientStorageGateway(db)}
}
