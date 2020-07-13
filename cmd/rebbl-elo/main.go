package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"

	"log"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
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
}
