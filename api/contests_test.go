package api

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	sample, _ := ioutil.ReadFile("testdata/getcontestssample.json")
	emptyJSON := []byte("{}")
	emptyContests := []byte("{\"contests\":[]}")

	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []*Contest
		wantErr bool
	}{
		{"Parse Sample Payload", args{sample}, wantedPayload.Contests, false},
		{"Error on Empty Payload", args{[]byte{}}, nil, true},
		{"Parse Empty JSON", args{emptyJSON}, nil, false},
		{"Parse Empty Contests List", args{emptyContests}, []*Contest{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseContests(tt.args.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

var wantedContest1 Contest = Contest{
	ContestID:       1071358,
	PlatformID:      1,
	CompetitionID:   191991,
	CompetitionName: "Season 14 - Division 4A",
	CurrentRound:    8,
	Format:          "round_robin",
	LeagueID:        42291,
	LeagueName:      "REBBL - GMan",
	MaxRound:        13,
	Status:          2,
	TeamAway: Team{
		TeamName:  "Rex Lupus",
		TeamID:    3213705,
		TeamLogo:  "Necromantic_10",
		TeamValue: 1560,
		Race:      "Necromantic",
		CoachName: "Jojishi",
		CoachID:   93409,
		Score:     1,
	},
	TeamHome: Team{
		TeamName:  "Face Planters",
		TeamID:    2988690,
		TeamLogo:  "HighElf_07",
		TeamValue: 1460,
		Race:      "HighElf",
		CoachName: "Spoon777",
		CoachID:   268730,
		Score:     2,
	},
	MatchUUID: []string{"10007e446c"},
}

var wantedContest2 Contest = Contest{
	ContestID:       1071360,
	PlatformID:      1,
	CompetitionID:   191991,
	CompetitionName: "Season 14 - Division 4A",
	CurrentRound:    8,
	Format:          "round_robin",
	LeagueID:        42291,
	LeagueName:      "REBBL - GMan",
	MaxRound:        13,
	Status:          2,
	TeamAway: Team{
		TeamName:  "Jellybean Jumparounds",
		TeamID:    3205862,
		TeamLogo:  "Elf_05",
		TeamValue: 1450,
		Race:      "ProElf",
		CoachName: "MumboJambo",
		CoachID:   163862,
		Score:     2,
	},
	TeamHome: Team{
		TeamName:  "Cursed Cowboy's",
		TeamID:    3031844,
		TeamLogo:  "Chaos_18",
		TeamValue: 1750,
		Race:      "Chaos",
		CoachName: "Dimmy Gee",
		CoachID:   305808,
		Score:     2,
	},
	MatchUUID: []string{"10007e4898"},
}

var wantedPayload ContestsPayload = ContestsPayload{
	Contests: []*Contest{&wantedContest1, &wantedContest2},
}
