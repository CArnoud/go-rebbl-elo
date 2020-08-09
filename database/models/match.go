package models

import (
	"github.com/CArnoud/go-rebbl-elo/api"

	"github.com/jinzhu/gorm"
)

// Match .
type Match struct {
	gorm.Model
	HomeTeamScore uint
	AwayTeamScore uint
	Started       string
	Finished      string
	Round         uint
	HomeTeamID    uint
	AwayTeamID    uint
	WinnerID      *uint
	CompetitionID uint
}

func getWinnerID(contest *api.Contest) *uint {
	var winnerID *uint = nil
	if contest.TeamAway.Score > contest.TeamHome.Score {
		temp := uint(contest.TeamAway.TeamID)
		winnerID = &temp
	} else {
		if contest.TeamHome.Score > contest.TeamAway.Score {
			temp := uint(contest.TeamHome.TeamID)
			winnerID = &temp
		}
	}

	return winnerID
}

// NewMatch returns a Match instance from a Contest instance.
func NewMatch(contest *api.Contest) *Match {
	return &Match{
		Model:         gorm.Model{ID: uint(contest.ContestID)},
		HomeTeamID:    uint(contest.TeamHome.TeamID),
		HomeTeamScore: uint(contest.TeamHome.Score),
		AwayTeamID:    uint(contest.TeamAway.TeamID),
		AwayTeamScore: uint(contest.TeamAway.Score),
		WinnerID:      getWinnerID(contest),
		Started:       contest.Started,
		Finished:      contest.Finished,
		Round:         uint(contest.CurrentRound),
		CompetitionID: uint(contest.CompetitionID),
	}
}
