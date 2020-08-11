package models

import (
	"github.com/jinzhu/gorm"
)

// Predictor is the internal representation of a predictor.
type Predictor struct {
	gorm.Model
	Name   string
	Config string
}
