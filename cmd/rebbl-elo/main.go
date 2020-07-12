package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	log.Println("connected")

	api.ParseMatch([]byte{})
}
