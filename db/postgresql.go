package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgreSQL() (*sql.DB, error) {
	 db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable")

	 if err != nil {
        fmt.Println("Error connecting to PostgreSQL database:", err)
        return nil, err
    }


		 // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        return nil, err
    }

		 return db, nil
}