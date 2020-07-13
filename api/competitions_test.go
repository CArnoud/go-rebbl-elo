package api

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseCompetitions(t *testing.T) {
	sample, _ := ioutil.ReadFile("testdata/getcompetitionssample.json")
	emptyJSON := []byte("{}")
	emptyCompetitions := []byte("{\"competitions\":[]}")

	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []*Competition
		wantErr bool
	}{
		{"Parse Sample Payload", args{sample}, wantedCompetitionsPayload.Competitions, false},
		{"Error on Empty Payload", args{[]byte{}}, nil, true},
		{"Parse Empty JSON", args{emptyJSON}, nil, false},
		{"Parse Empty Competitions List", args{emptyCompetitions}, []*Competition{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCompetitions(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCompetitions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCompetitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

var wantedCompetition1 = Competition{
	ID: 191991,
	PlatformID: 1, 
	DateCreated: "2020-05-17 14:36:49",
	Format: "round_robin",
	LeagueID: 42291,
	Name: "Season 14 - Division 4A",    
	Round: 8,
	RoundsCount: 13,
	Status: 1,
	TeamsCount: 14,   
	TeamsMax: 14,     
	TurnDuration: 3,
}

var wantedCompetition2 = Competition{
	ID: 191988,
	PlatformID: 1, 
	DateCreated: "2020-05-17 14:36:49",
	Format: "round_robin",
	LeagueID: 42291,
	Name: "Season 14 - Division 3A",    
	Round: 8,
	RoundsCount: 13,
	Status: 1,
	TeamsCount: 14,   
	TeamsMax: 14,     
	TurnDuration: 3,
}

var wantedCompetitionsPayload CompetitionsPayload = CompetitionsPayload{
	Competitions: []*Competition{&wantedCompetition1, &wantedCompetition2},
}

