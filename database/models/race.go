package models

import (
	"github.com/jinzhu/gorm"
)

// Race represent a team's race
type Race struct {
	gorm.Model
	Name string
	Code uint

	Teams []*Team
}
