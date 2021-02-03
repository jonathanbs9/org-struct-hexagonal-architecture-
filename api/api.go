package api

import (
	"github.com/org-struct-hexagonal-architecture-/internal/storage"
	patients "github.com/org-struct-hexagonal-architecture-/patients/web_application"
)

// Start func => va a ser llamado por el main para arrancar la app
func Start(port string) {
	db := storage.ConnecToDB()
	defer db.Close()

	r := routes(patients.NewPatientHTTPService(db))
	server := newServer(port, r)

	server.Start()
}
