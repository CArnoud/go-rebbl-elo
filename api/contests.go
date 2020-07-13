package api

import (
	"encoding/json"
)

// ParseContests .
func ParseContests(payload []byte) ([]*Contest, error) {
	var result ContestsPayload

	err := json.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	return result.Contests, nil
}

// ContestsPayload - complete payload from the contests API request.
type ContestsPayload struct {
	Contests []*Contest `json:"contests"`
}

// Contest - JSON representing a match from Spike.
type Contest struct {
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
}

// Team - JSON representation of a Team inside a Contest.
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
