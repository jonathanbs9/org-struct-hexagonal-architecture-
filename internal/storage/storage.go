package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// ConnecToDB func
func ConnecToDB() *sql.DB {
	server := os.Getenv("GCP_HOSTNAME")
	if server == "" {
		server = "localhost"
	}

	connURL := fmt.Sprintf("postgres://postgres:postgres@%s/project?sslmode=disable", server)
	db, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Fatalf("Fallo la conexión  a la base de datos via %s :  % v ", connURL, err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Falló el ping a la base de datos via %s : %v ", connURL, err)
	}
	log.Println("Conexión exitosa")
	return db

}
