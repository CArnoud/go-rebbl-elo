package models

import (
	"github.com/jinzhu/gorm"
)

// Team holds information about a REBBL team
type Team struct {
	gorm.Model
	Name    string
	RebblID string
}
