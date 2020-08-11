package predictors

import (
	"github.com/CArnoud/go-rebbl-elo/database/models"
)

// HomeTeamPicker always picks the home team.
type HomeTeamPicker struct {
	Draws uint
	CorrectPicks uint
	TotalGames uint
}

// NewHomeTeamPicker instantiates HomeTeamPicker.
func NewHomeTeamPicker() *HomeTeamPicker {
	return &HomeTeamPicker{}
}

// ProcessMatch updates ratings for both teams and checks whether prediction was correct.
func (htp *HomeTeamPicker) ProcessMatch(match *models.Match) {
	htp.TotalGames++

	if match.WinnerID == nil {
		htp.Draws++
	} else {
		if *match.WinnerID == match.HomeTeamID {
			htp.CorrectPicks++
		}
	}
}