package api

import (
	"encoding/json"
)

// ParseMatch .
func ParseMatch(payload []byte) (*Match, error) {
	var match Match

	err := json.Unmarshal(payload, &match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}

// Match - JSON representing a match from Spike
type Match struct {
	ID              string   `json:"_id"`
	ContestID       int      `json:"contest_id"`
	PlatformID      int      `json:"platform_id"`
	CompetitionID   int      `json:"competition_id"`
	CompetitionLogo string   `json:"competition_logo"`
	CompetitionName string   `json:"competition_name"`
	CurrentRound    int      `json:"current_round"`
	Format          string   `json:"format"`
	LastUpdate      string   `json:"last_update"`
	LeagueID        int      `json:"league_id"`
	LeagueName      string   `json:"league_name"`
	MaxRound        int      `json:"max_round"`
	Status          int      `json:"status"`
	TeamAway        Team     `json:"team_away"`
	TeamHome        Team     `json:"team_home"`
	MatchUUID       []string `json:"match_uuid"`
	Winner          Winner   `json:"winner"`
}

// Team JSON representation of a Team inside a Match
type Team struct {
	TeamName  string `json:"team_name"`
	TeamID    int    `json:"team_id"`
	TeamLogo  string `json:"team_logo"`
	TeamValue int    `json:"team_value"`
	Race      string `json:"race"`
	CoachName string `json:"coach_name"`
	CoachID   int    `json:"coach_id"`
	Score     int    `json:"score"`
}

// Winner JSON indicating who won a match
type Winner struct {
	CoachID int `json:"coach_id"`
	TeamID  int `json:"team_id"`
}
