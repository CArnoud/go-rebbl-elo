package models

import (
	"github.com/jinzhu/gorm"
)

// Season represents a group of competitions which happened on the same REBBL season
type Season struct {
	gorm.Model
	Name  string
	Index uint

	Competitions []*Competition
}
