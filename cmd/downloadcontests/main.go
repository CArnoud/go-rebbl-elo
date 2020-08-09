package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database"
	"github.com/CArnoud/go-rebbl-elo/database/models"

	"github.com/CArnoud/go-rebbl-elo/download"

	"log"
	"net/http"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	spikeClient := api.NewSpikeClient(cfg, http.DefaultClient)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Unable to open database connection: " + err.Error())
	}

	downloader := download.NewDownloader(spikeClient, db)

	competitionRows, err := db.RawFind("competitions", "id, name, league_id")
	if err != nil {
		log.Fatal(err)
	}
	defer competitionRows.Close()

	for competitionRows.Next() {
		// TODO goroutine?
		var id uint
		var name string
		var leagueID uint
		competitionRows.Scan(&id, &name, &leagueID)

		if models.IsCompetitive(name) || leagueID == 34500 {
			err := downloader.DownloadContests(id)
			if err != nil {
				log.Printf("%s (%d) Error: %s", name, id, err)
			}

			log.Printf("%s (%d) success.", name, id)
		} else {
			log.Printf("%s (%d) ignored.", name, id)
		}
	}
}
