package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func OpenPostgresDB() (db *gorm.DB, err error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	//user := "postgres"
	//password := "123456"
	//host := "localhost"
	//port := "5432"
	//dbname := "gocraw"

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		err = fmt.Errorf("[infra.PostgresOpen] failed to open connection to database: %w", err)
	}

	return
}
