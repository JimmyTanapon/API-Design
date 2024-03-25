package user

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	userd    = "postgres"   // as defined in docker-compose.yml
	password = "jumonji123" // as defined in docker-compose.yml
	dbname   = "users"      // as defined in docker-compose.yml
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userd, password, dbname)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return sdb
}
