package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database"
	"github.com/CArnoud/go-rebbl-elo/download"

	"log"
	"net/http"
)

func init() {
	log.SetFlags(0)
}

func main() {
	leagueList := []uint{
		42290, // REL
		42291, // GMAN
		42292, // BIGO
		75681, // REL2
		75684, // GMAN2
		34500, // Playoffs
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Unable to open database connection: " + err.Error())
	}

	spikeClient := api.NewSpikeClient(cfg, http.DefaultClient)
	downloader := download.NewDownloader(spikeClient, db)

	for _, leagueID := range leagueList {
		err := downloader.DownloadCompetitions(leagueID)
		if err != nil {
			log.Printf("League %d download error: %s", leagueID, err.Error())
		} else {
			log.Printf("League %d download success.", leagueID)
		}
	}
}
