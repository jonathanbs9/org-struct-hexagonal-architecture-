package patients

import "time"

// Patient struct => Estructura de un paciente
type Patient struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// CreatePatientCMD struct => Estructura para crear un paciente
type CreatePatientCMD struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
