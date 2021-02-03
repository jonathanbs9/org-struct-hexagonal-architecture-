package api

import (
	"github.com/go-chi/chi"
	patients "github.com/org-struct-hexagonal-architecture-/patients/web_application"
)

func routes(services *patients.PatientHTTPService) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", services.GetPatientsHandler)
	r.Post("/patients", services.CreatePatientsHandler)
	r.Get("/patients/{id}", services.GetPatientsByIDHandler)

	return r
}
