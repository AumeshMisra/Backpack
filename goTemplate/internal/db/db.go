package db

import (
	"log"
    "os"
    "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
    var hostName string = "localhost" 
    if (os.Getenv("DB_HOST") == "docker") {
        hostName = "db"
    }
    var dbURL string = fmt.Sprintf("postgres://postgres:postgres@%s:5432/db", hostName)
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    return db
}