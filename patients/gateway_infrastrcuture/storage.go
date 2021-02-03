package patients

import (
	"database/sql"
	"log"
	"time"

	patients "github.com/org-struct-hexagonal-architecture-/patients/models_domain"
)

// PatientStorage interface
type PatientStorage interface {
	createPatientDB(p *patients.CreatePatientCMD) (*patients.Patient, error)
	getPatientsDB() []*patients.Patient
	getPatientByIDDB(id int64) (*patients.Patient, error)
}

// PatientService struct
type PatientService struct {
	db *sql.DB
}

// NewPatientStorageGateway func
func NewPatientStorageGateway(db *sql.DB) PatientStorage {
	return &PatientService{db: db}
}

// CreatePatientDB func
func (s *PatientService) createPatientDB(p *patients.CreatePatientCMD) (*patients.Patient, error) {
	log.Println("Creando nuevo paciente")
	res, err := s.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?, ?, ?, ?, ?)", p.FirstName, p.LastName, p.Address, p.Phone, p.Email)

	if err != nil {
		log.Println("No se puede crear paciente en DB | " + err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &patients.Patient{
		ID:        id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Address:   p.Address,
		Phone:     p.Phone,
		Email:     p.Email,
		CreatedAt: time.Now(),
	}, nil
}

func (s *PatientService) getPatientsDB() []*patients.Patient {
	rows, err := s.db.Query("select id, first_name, last_name, address, phone, email, created_at from patient")

	if err != nil {
		log.Printf("No se puede ejecutar la consulta select => " + err.Error())
		return nil
	}

	defer rows.Close()
	var p []*patients.Patient
	for rows.Next() {
		var patient patients.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone, &patient.Email, &patient.CreatedAt)
		if err != nil {
			log.Printf("No se puede leer la fila actual")
			return nil
		}
		p = append(p, &patient)
	}
	return p
}

func (s *PatientService) getPatientByIDDB(id int64) (*patients.Patient, error) {
	var patient patients.Patient
	err := s.db.QueryRow(`select id, first_name, last_name, address, phone, email, created_at from patient
		where id = ?`, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone, &patient.Email, &patient.CreatedAt)

	if err != nil {
		log.Printf("No se puede encontrar paciente")
		return nil, err
	}
	return &patient, nil
}
