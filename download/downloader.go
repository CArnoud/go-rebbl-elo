package download

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/CArnoud/go-rebbl-elo/database/models"

	"database/sql"
	"fmt"
	"log"
	"sync"
)

// APIClient represents the interaction with an external API to get contests.
type APIClient interface {
	GetContests(uint, uint) ([]byte, error)
	GetCompetitions(uint, string) ([]byte, error)
}

// DBClient represents the interaction with a database management system.
type DBClient interface {
	FirstOrCreate(interface{}, ...interface{}) error
	RawFind(string, string, string) (*sql.Rows, error)
	Delete(interface{}) error
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

// DownloadCompetitions downloads all competitions from a league.
func (d *Downloader) DownloadCompetitions(leagueID uint) error {
	resp, err := d.api.GetCompetitions(leagueID, "1")
	if err != nil {
		return err
	}

	resp2, err := d.api.GetCompetitions(leagueID, "2")
	if err != nil {
		return err
	}

	apiCompetitions, err := api.ParseCompetitions(resp)
	if err != nil {
		return err
	}

	temp, err := api.ParseCompetitions(resp2)
	if err != nil {
		return err
	}

	apiCompetitions = append(apiCompetitions, temp...)

	for _, apiCompetition := range apiCompetitions {
		c := models.NewCompetition(apiCompetition)
		if models.IsCompetitionValid(c) {
			err := d.db.FirstOrCreate(c)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Ignoring %s", c.Name)
		}
	}

	return nil
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

	// TODO in-progress (status 1)

	contests := append(played, scheduled...)

	var wg sync.WaitGroup
	for _, contest := range contests {
		wg.Add(1)
		go d.insertContestWorker(contest, &wg)
	}

	wg.Wait()

	return nil
}

func (d *Downloader) insertContestWorker(contest *api.Contest, wg *sync.WaitGroup) {
	defer wg.Done()

	if api.ContestIsCompetitive(contest) {
		errs := []error{}

		homeCoach := models.NewCoach(&contest.TeamHome)
		homeCoachIDString := fmt.Sprint(homeCoach.ID)
		errs = append(errs, d.db.FirstOrCreate(&homeCoach, homeCoachIDString))

		awayCoach := models.NewCoach(&contest.TeamAway)
		awayCoachIDString := fmt.Sprint(awayCoach.ID)
		errs = append(errs, d.db.FirstOrCreate(&awayCoach, awayCoachIDString))

		homeTeam := models.NewTeam(&contest.TeamHome)
		homeTeamIDString := fmt.Sprint(homeTeam.ID)
		errs = append(errs, d.db.FirstOrCreate(&homeTeam, homeTeamIDString))

		awayTeam := models.NewTeam(&contest.TeamAway)
		awayTeamIDString := fmt.Sprint(awayTeam.ID)
		errs = append(errs, d.db.FirstOrCreate(&awayTeam, awayTeamIDString))

		if errs[0] != nil || errs[1] != nil || errs[2] != nil || errs[3] != nil {
			log.Printf("Competition %d errors adding dependencies: %s", contest.CompetitionID, errs)
		}

		model := models.NewMatch(contest)
		if contest.Finished != "" {
			err := d.db.Delete(&model)
			if err != nil {
				log.Printf("*** Unable to delete contest %d: %s", contest.ContestID, err.Error())
			}
		}

		modelIDString := fmt.Sprint(model.ID)
		d.db.FirstOrCreate(&model, modelIDString)
	}
}
