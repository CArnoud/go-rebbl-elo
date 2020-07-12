package models

import (
	"github.com/jinzhu/gorm"
)

// Competition belongs to a League and a Season
type Competition struct {
	gorm.Model
	Name     string
	LeagueID uint
	SeasonID uint

	Rounds []*Round
}
