package models

import (
	"github.com/jinzhu/gorm"
)

// Rating represents a team's rating according to a predictor.
type Rating struct {
	gorm.Model
	PredictorID uint
	TeamID      uint
	MatchID     uint
	Value       int
}
