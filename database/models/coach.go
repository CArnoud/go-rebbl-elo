package models

import (
	"github.com/jinzhu/gorm"
)

// Coach has teams
type Coach struct {
	gorm.Model
	Name    string
	RebblID string

	Teams []*Team
}
