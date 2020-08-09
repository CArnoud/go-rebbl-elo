package main

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database"
	"github.com/CArnoud/go-rebbl-elo/download"

	"database/sql"
	"log"
	"net/http"
	"sync"
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
	defer competitionRows.Close()
	if err != nil {
		log.Fatal(err)
	}	

	var wg sync.WaitGroup
	for competitionRows.Next() {
		wg.Add(1)
		go processCompetition(downloader, competitionRows, &wg)
	}

	wg.Wait()
}

func processCompetition(downloader *download.Downloader, competitionRows *sql.Rows, wg *sync.WaitGroup) {
	defer wg.Done()
	var id uint
	var name string
	var leagueID uint
	competitionRows.Scan(&id, &name, &leagueID)

	if id != 0 && name != "" {
		err := downloader.DownloadContests(id)
		if err != nil {
			log.Printf("%s (%d) Error: %s", name, id, err)
		} else {
			log.Printf("%s (%d) success.", name, id)
		}
	} else {
		log.Printf("Ignored Row: (%d, %s, %d)", id, name, leagueID)
	}
}
