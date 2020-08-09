package download

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/database/models"

	"database/sql"
	"fmt"
	"log"
)

// APIClient represents the interaction with an external API to get contests.
type APIClient interface {
	GetContests(uint, uint) ([]byte, error)
}

// DBClient represents the interaction with a database management system.
type DBClient interface {
	FirstOrCreate(interface{}, ...interface{}) error
	RawFind(string, string) (*sql.Rows, error)
}

// Downloader manages the download and insertion of records.
type Downloader struct {
	api APIClient
	db  DBClient
}

// NewDownloader instantiates a Downloader given an API client and a database client.
func NewDownloader(api APIClient, db DBClient) *Downloader {
	return &Downloader{
		api: api,
		db:  db,
	}
}

// DownloadContests downloads all contests in a competition.
func (d *Downloader) DownloadContests(competitionID uint) error {
	var played []*api.Contest
	var scheduled []*api.Contest

	resp, err := d.api.GetContests(competitionID, 2)
	if err == nil {
		played, err = api.ParseContests(resp)
		if err != nil {
			return err
		}
	} else {
		log.Printf("Competition %d error: %s", competitionID, err.Error())
	}

	resp2, err := d.api.GetContests(competitionID, 0)
	if err == nil {
		scheduled, err = api.ParseContests(resp2)
		if err != nil {
			return err
		}
	} else {
		log.Printf("Competition %d error: %s", competitionID, err.Error())
	}

	contests := append(played, scheduled...)

	for _, contest := range contests {
		if api.ContestIsCompetitive(contest) {
			homeCoach := models.NewCoach(&contest.TeamHome)
			homeCoachIDString := fmt.Sprint(homeCoach.ID)
			d.db.FirstOrCreate(&homeCoach, homeCoachIDString)

			awayCoach := models.NewCoach(&contest.TeamAway)
			awayCoachIDString := fmt.Sprint(awayCoach.ID)
			d.db.FirstOrCreate(&awayCoach, awayCoachIDString)

			homeTeam := models.NewTeam(&contest.TeamHome)
			homeTeamIDString := fmt.Sprint(homeTeam.ID)
			d.db.FirstOrCreate(&homeTeam, homeTeamIDString)

			awayTeam := models.NewTeam(&contest.TeamAway)
			awayTeamIDString := fmt.Sprint(awayTeam.ID)
			d.db.FirstOrCreate(&awayTeam, awayTeamIDString)

			// TODO upsert
			model := models.NewMatch(contest)
			modelIDString := fmt.Sprint(model.ID)
			d.db.FirstOrCreate(&model, modelIDString)
		}
	}

	return nil
}
