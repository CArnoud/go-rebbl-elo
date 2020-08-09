package models

import (
	"github.com/CArnoud/go-rebbl-elo/api"

	"github.com/jinzhu/gorm"
)

// Team holds information about a REBBL team
type Team struct {
	gorm.Model
	Name    string
	Logo    string
	CoachID uint
	RaceID  uint
}

// NewTeam creates a Team from Spike API payload.
func NewTeam(t *api.Team) *Team {
	return &Team{
		Model:   gorm.Model{ID: uint(t.TeamID)},
		Name:    t.TeamName,
		Logo:    t.TeamLogo,
		CoachID: uint(t.CoachID),
		RaceID:  1, // TODO: change to real race
	}
}
