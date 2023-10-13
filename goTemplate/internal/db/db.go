package db

import (
	"log"
    "os"
    "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
    var db_host string = "localhost";
    var db_username = os.Getenv("DB_USER");
    var db_password = os.Getenv("DB_PASSWORD");
    var db_port = os.Getenv("DB_PORT");
    var db_name = os.Getenv("DB_NAME")
    if (os.Getenv("GO_ENV") == "docker") {
        db_host = "db"
    }
    var dbURL string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_username, db_password, db_host, db_port, db_name)
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    return db
}