// config/db.go
package config

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v4"
)

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	fmt.Println("Connected to PostgreSQL on port 5433")
	return conn
}
