package models

import (
	"github.com/jinzhu/gorm"
)

// Round belongs to a competition
type Round struct {
	gorm.Model
	Index uint

	Matches []*Match
}
