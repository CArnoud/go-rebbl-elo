package database

import (
	"github.com/CArnoud/go-rebbl-elo/config"
	"github.com/CArnoud/go-rebbl-elo/database/models"

	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // tells gorm which driver to use
)

// SQLClient .
type SQLClient interface {
	AutoMigrate(dst ...interface{}) *gorm.DB
	DropTableIfExists(dst ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB
	Raw(table string, args ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}

// Database .
type Database struct {
	client SQLClient
}

func (d *Database) addforeignKeys(db *gorm.DB) {
	db.Model(&models.Team{}).AddForeignKey("coach_id", "coaches(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Team{}).AddForeignKey("race_id", "races(id)", "RESTRICT", "RESTRICT")

	db.Model(&models.Match{}).AddForeignKey("competition_id", "competitions(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Match{}).AddForeignKey("home_team_id", "teams(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Match{}).AddForeignKey("away_team_id", "teams(id)", "RESTRICT", "RESTRICT")

	db.Model(&models.Competition{}).AddForeignKey("league_id", "leagues(id)", "RESTRICT", "RESTRICT")

	db.Model(&models.Rating{}).AddForeignKey("predictor_id", "predictors(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Rating{}).AddForeignKey("team_id", "teams(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Rating{}).AddForeignKey("match_id", "matches(id)", "RESTRICT", "RESTRICT")
}

// AutoMigrate .
func (d *Database) AutoMigrate() {
	d.client.DropTableIfExists(
		&models.Rating{},
		&models.Predictor{},
		&models.Match{},
		&models.Competition{},
		&models.League{},
		&models.Team{},
		&models.Coach{},
		&models.Race{},
	)

	gormDB := d.client.AutoMigrate(
		&models.League{},
		&models.Coach{},
		&models.Competition{},
		&models.Match{},
		&models.Race{},
		&models.Team{},
		&models.Predictor{},
		&models.Rating{},
	)

	d.addforeignKeys(gormDB)
}

// Create .
func (d *Database) Create(value interface{}) error {
	return d.client.Create(value).Error
}

// FirstOrCreate .
func (d *Database) FirstOrCreate(out interface{}, where ...interface{}) error {
	return d.client.FirstOrCreate(out, where).Error
}

// RawFind .
func (d *Database) RawFind(table string, columns string, orderBy string) (*sql.Rows, error) {
	queryString := fmt.Sprintf("SELECT %s FROM public.%s;", columns, table)
	if orderBy != "" {
		queryString = fmt.Sprintf("SELECT %s FROM public.%s ORDER BY %s;", columns, table, orderBy)
	}
	return d.client.Raw(queryString).Rows()
}

// Delete .
func (d *Database) Delete(value interface{}) error {
	return d.client.Delete(value).Error
}

// NewDatabase instantiates a database client.
func NewDatabase(c *config.Config) (*Database, error) {
	options := "host=" + c.DatabaseHost + " port=" + c.DatabasePort + " user=" + c.DatabaseUser + " dbname=" + c.DatabaseName + " password=" + c.DatabasePassword + " sslmode=disable"
	db, err := gorm.Open("postgres", options)
	if err != nil {
		return nil, err
	}
	db.LogMode(false)

	return &Database{db}, nil
}
