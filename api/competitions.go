package api

import (
	"encoding/json"
)

// ParseCompetitions .
func ParseCompetitions(payload []byte) ([]*Competition, error) {
	var result CompetitionsPayload

	err := json.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	return result.Competitions, nil
}

// CompetitionsPayload - complete payload from Spike competitions endpoint.
type CompetitionsPayload struct {
	Competitions []*Competition `json:"competitions"`
}

// Competition - JSON representation of a competition on the Spike website.
type Competition struct {
	ID           int    `json:"id"`
	PlatformID   int    `json:"platform_id"`
	DateCreated  string `json:"date_created"`
	Format       string `json:"format"`
	LeagueID     int    `json:"league_id"`
	Name         string `json:"name"`
	Round        int    `json:"round"`
	RoundsCount  int    `json:"rounds_count"`
	Status       int    `json:"status"`
	TeamsCount   int    `json:"teams_count"`
	TeamsMax     int    `json:"teams_max"`
	TurnDuration int    `json:"turn_duration"`
}