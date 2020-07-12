package api

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseMatch(t *testing.T) {
	sample, _ := ioutil.ReadFile("testdata/matchsample.json")

	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Match
		wantErr bool
	}{
		{"Parse sample match JSON", args{sample}, &wantedMatch, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMatch(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

var wantedMatch Match = Match{
	ID:              "5ec5be341c4e587853fb59cf",
	ContestID:       1070332,
	PlatformID:      1,
	CompetitionID:   191985,
	CompetitionLogo: ":Logo_Neutre_23:",
	CompetitionName: "Season 14 - Division 1",
	CurrentRound:    8,
	Format:          "round_robin",
	LastUpdate:      "2020-07-12T02:04:40.235173",
	LeagueID:        42291,
	LeagueName:      "REBBL - GMan",
	MaxRound:        13,
	Status:          2,
	TeamAway: Team{
		TeamName:  "New Yorc Pilanders",
		TeamID:    2975142,
		TeamLogo:  "Orc_01",
		TeamValue: 2100,
		Race:      "Orc",
		CoachName: "Largo Feldlauf",
		CoachID:   192942,
		Score:     1,
	},
	TeamHome: Team{
		TeamName:  "BreakNec Rampage",
		TeamID:    2560533,
		TeamLogo:  "Necromantic_03",
		TeamValue: 1400,
		Race:      "Necromantic",
		CoachName: "O'Kim",
		CoachID:   13627,
		Score:     0,
	},
	MatchUUID: []string{"10007e0676"},
	Winner: Winner{
		CoachID: 192942,
		TeamID:  2975142,
	},
}
