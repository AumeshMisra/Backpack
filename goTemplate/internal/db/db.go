package db

import (
	"log"
    "os"
    "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
    db_host := "localhost"
    db_username := os.Getenv("DB_USER")
    db_password := os.Getenv("DB_PASSWORD")
    db_port := os.Getenv("DB_PORT")
    db_name := os.Getenv("DB_NAME")
    if os.Getenv("GO_ENV") == "docker" {
        db_host = "db"
    }
    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_username, db_password, db_host, db_port, db_name)
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    return db
}