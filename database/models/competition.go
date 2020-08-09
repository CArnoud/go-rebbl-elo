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

// IsCompetitive evaluates the name of a competition to determined if it should be ignored.
func IsCompetitive(name string) bool {
	return (strings.Contains(name, " Division") ||
		strings.Contains(name, " Swiss") ||
		strings.Contains(name, "Season12 - Div 3") ||
		strings.Contains(name, " Play-Ins") ||
		strings.Contains(name, " Challenger's Cup") ||
		strings.Contains(name, " Playoffs")) &&
		!strings.Contains(strings.ToLower(name), " mng")
}
