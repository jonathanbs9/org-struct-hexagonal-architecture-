package patients

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/org-struct-hexagonal-architecture-/internal/web"
	patients "github.com/org-struct-hexagonal-architecture-/patients/gateway_infrastrcuture"
	models "github.com/org-struct-hexagonal-architecture-/patients/models_domain"
)

// PatientHTTPService struct
type PatientHTTPService struct {
	gtw patients.PatientGateway
}

// NewPatientHTTPService  func
func NewPatientHTTPService(db *sql.DB) *PatientHTTPService {
	return &PatientHTTPService{
		gtw: patients.NewPatientGateway(db),
	}
}

// GetPatientsHandler func
func (s *PatientHTTPService) GetPatientsHandler(w http.ResponseWriter, r *http.Request) {
	p := s.gtw.GetPatients()
	if p == nill || len(p) == 0 {
		p = []*models.Patient{}
	}
	web.Success(&p, http.StatusOK).Send(w)
}

// GetPatientsByIDHandler func
func (s *PatientHTTPService) GetPatientsByIDHandler(w http.ResponseWriter, r *http.Request) {
	patientID := chi.URLParam(r, "patientID")
	id, _ := strconv.ParseInt(patientID, 10, 64)
	patient, err := s.gtw.GetPatientByID(id)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&patient, http.StatusOK).Send(w)
}

// CreatePatientsHandler func
func (s *PatientHTTPService) CreatePatientsHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	var cmd models.CreatePatientCMD

}
