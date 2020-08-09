package models

import (
	"github.com/CArnoud/go-rebbl-elo/api"
	"github.com/jinzhu/gorm"
)

// Coach has teams
type Coach struct {
	gorm.Model
	Name  string
	Teams []*Team
}

// NewCoach creates a Coach from Spike API team.
func NewCoach(t *api.Team) *Coach {
	return &Coach{
		Model: gorm.Model{ID: uint(t.CoachID)},
		Name:  t.CoachName,
	}
}
