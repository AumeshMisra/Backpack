package main

import (
	"goTemplate/internal/db"
	"log"
)

func main() {
	DB := db.Init()
	log.Println(DB)
}