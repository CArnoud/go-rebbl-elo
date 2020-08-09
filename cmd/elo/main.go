package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"
	// "github.com/CArnoud/go-rebbl-elo/database"

	"log"
	"net/http"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database: " + err.Error())
	// }
	// defer db.Close()

	// log.Println("connected")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	spikeClient := api.NewSpikeClient(cfg, http.DefaultClient)
	// resp, err := spikeClient.GetCompetitions(42291)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	resp, err := spikeClient.GetContests(191991, 2)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(resp))

	// db, err := database.NewDatabase()
	// if err != nil {
	// 	log.Fatal("Unable to open database connection: " + err.Error())
	// }

	// db.AutoMigrate()
}
