package main

import (
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database"
	"github.com/CArnoud/go-rebbl-elo/database/models"
	"github.com/CArnoud/go-rebbl-elo/predictors"
	"github.com/jinzhu/gorm"

	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Unable to open database connection: " + err.Error())
	}

	wantedColumns := "id, home_team_score, away_team_score, home_team_id, away_team_id, winner_id, finished"
	matchRows, err := db.RawFind("matches", wantedColumns, "finished")
	if err != nil {
		log.Fatalf("Unable to select matches: %s", err.Error())
	}

	predictor := predictors.NewTeamElo(100, 700, 1000)
	// predictor := predictors.NewHomeTeamPicker()

	for matchRows.Next() {
		var id uint
		var homeTeamScore uint
		var awayTeamScore uint
		var homeTeamID uint
		var awayTeamID uint
		var winnerID *uint
		var finished string

		err := matchRows.Scan(&id, &homeTeamScore, &awayTeamScore, &homeTeamID, &awayTeamID, &winnerID, &finished)
		if err != nil {
			log.Printf("Unable to scan rows: %s", err.Error())
		} else {
			if finished != "" {
				match := models.Match{
					Model: gorm.Model{ID: id},
					HomeTeamScore: homeTeamScore,
					AwayTeamScore: awayTeamScore,
					HomeTeamID: homeTeamID,
					AwayTeamID: awayTeamID,
					WinnerID: winnerID,
				}
				predictor.ProcessMatch(&match)
			} else {
				log.Printf("Match %d unplayed", id)
			}
		}
	}

	log.Printf("Correct picks: %d out of %d games (%d draws)", predictor.CorrectPicks, predictor.TotalGames, predictor.Draws)
}
