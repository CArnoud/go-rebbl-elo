package models

import (
	"github.com/jinzhu/gorm"
)

// Match belongs to a Round
type Match struct {
	gorm.Model
	RebblID    string
	HomeTeamID uint
	AwayTeamID uint
	WinnerID   uint
}
