package predictors

import (
	"log"

	"github.com/CArnoud/go-rebbl-elo/database/models"

	"github.com/kortemy/elo-go"
)

// TeamElo predictor uses team results to make ratings and predictions.
type TeamElo struct {
	elo *elogo.Elo
	ratings map[uint]*int
	norm int
	TotalGames uint
	CorrectPicks uint
	Draws uint
}

// NewTeamElo initializes an elo predictor based on team information.
func NewTeamElo(kFactor int, deviation int, norm int) *TeamElo {
	return &TeamElo{
		elo: elogo.NewEloWithFactors(kFactor, deviation),
		ratings: make(map[uint]*int),
		norm: norm,
	}
}

func (te *TeamElo) getRating(id uint) int {
	if te.ratings[id] == nil {
		te.ratings[id] = &te.norm
	}

	return *te.ratings[id]
}

// ProcessMatch updates ratings for both teams and checks whether prediction was correct.
func (te *TeamElo) ProcessMatch(match *models.Match) {
	homeRating := te.getRating(match.HomeTeamID)
	awayRating := te.getRating(match.AwayTeamID)

	var homeScore float64
	if match.WinnerID == nil {
		homeScore = 0.5

		if homeRating != te.norm || awayRating != te.norm {
			te.Draws++
		}
	} else {
		if *match.WinnerID == match.HomeTeamID {
			homeScore = 1

			if homeRating > awayRating {
				te.CorrectPicks++
			}
		} else {
			homeScore = -1

			if awayRating > homeRating {
				te.CorrectPicks++
			}
		}
	}

	if homeRating != te.norm || awayRating != te.norm {
		te.TotalGames++
	}

	homeOutcome, awayOutcome := te.elo.Outcome(homeRating, awayRating, homeScore)
	te.ratings[match.HomeTeamID] = &homeOutcome.Rating
	te.ratings[match.AwayTeamID] = &awayOutcome.Rating

	log.Printf("team %d (%d) %d VS %d team %d (%d) => %d (%d), %d (%d)", match.HomeTeamID, homeRating, match.HomeTeamScore, match.AwayTeamScore, match.AwayTeamID, awayRating, homeOutcome.Rating, homeOutcome.Delta, awayOutcome.Rating, awayOutcome.Delta)
}

// ExpectedResult returns expected result of a team vs its opponent
func (te *TeamElo) ExpectedResult(teamRating int, opponentRating int) float64 {
	return te.elo.ExpectedScore(teamRating, opponentRating)
}
