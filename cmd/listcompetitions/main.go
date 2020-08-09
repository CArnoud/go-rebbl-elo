package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database"
	"github.com/CArnoud/go-rebbl-elo/database/models"

	// "fmt"
	"log"
	"net/http"
	// "github.com/jinzhu/gorm"
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

	for _, leagueID := range leagueList {
		spikeClient := api.NewSpikeClient(cfg, http.DefaultClient)
		resp, err := spikeClient.GetCompetitions(leagueID, "1")
		if err != nil {
			log.Fatal(err.Error())
		}

		resp2, err := spikeClient.GetCompetitions(leagueID, "2")
		if err != nil {
			log.Fatal(err.Error())
		}

		apiCompetitions, err := api.ParseCompetitions(resp)
		if err != nil {
			log.Fatal(err.Error())
		}

		temp, err := api.ParseCompetitions(resp2)
		if err != nil {
			log.Fatal(err.Error())
		}

		apiCompetitions = append(apiCompetitions, temp...)

		for _, apiCompetition := range apiCompetitions {
			c := models.NewCompetition(apiCompetition)
			db.FirstOrCreate(c)
		}
	}
}
