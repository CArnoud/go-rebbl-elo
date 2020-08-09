package models

import (
	"strings"

	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/jinzhu/gorm"
)

// Competition belongs to a League and a Season.
type Competition struct {
	gorm.Model
	Name         string
	Status       uint
	DateCreated  string
	Format       string
	RoundsCount  int
	TeamsCount   uint
	TeamsMax     uint
	TurnDuration uint
	LeagueID     uint
}

// NewCompetition creates a Competition instance from a Spike API payload
func NewCompetition(c *api.Competition) *Competition {
	return &Competition{
		Model:        gorm.Model{ID: uint(c.ID)},
		Name:         c.Name,
		Status:       uint(c.Status),
		DateCreated:  c.DateCreated,
		Format:       c.Format,
		RoundsCount:  c.RoundsCount,
		TeamsCount:   uint(c.TeamsCount),
		TeamsMax:     uint(c.TeamsMax),
		TurnDuration: uint(c.TurnDuration),
		LeagueID:     uint(c.LeagueID),
	}
}

// IsCompetitionValid evaluates the name of a competition to determine whether it should be ignored.
func IsCompetitionValid(c *Competition) bool {
	return strings.Contains(c.Name, " Division") ||
		strings.Contains(c.Name, " Swiss") ||
		strings.Contains(c.Name, "Season12 - Div 3") ||
		strings.Contains(c.Name, " Play-Ins") ||
		strings.Contains(c.Name, " Challenger's Cup") ||
		strings.Contains(c.Name, " Playoffs") ||
		(c.LeagueID == 34500 &&
			!strings.Contains(strings.ToLower(c.Name), "mng"))
}
