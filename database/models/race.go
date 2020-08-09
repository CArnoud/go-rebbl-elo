package models

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/jinzhu/gorm"
)

// Race represent a team's race
type Race struct {
	gorm.Model
	Name  string
	Code  uint
	Teams []*Team
}

// NewRace creates a race from a Spike API team
func NewRace(t *api.Team) *Race {
	return &Race{
		Name: t.Race,
	}
}
