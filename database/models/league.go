package models

import (
	"github.com/jinzhu/gorm"
)

// League values are REL, GMAN, BIGO
type League struct {
	gorm.Model
	Name         string
	Competitions []*Competition
}
