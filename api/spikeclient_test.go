package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/CArnoud/go-rebbl-elo/config"
)

var mockConfig = config.Config{}
var errMock = errors.New("mockerror")

type MockDoer struct {
	err        *error
	statusCode int
}

func (g *MockDoer) Do(req *http.Request) (*http.Response, error) {
	if g.err != nil {
		return nil, *g.err
	}

	body := ioutil.NopCloser(bytes.NewBufferString(req.URL.String()))
	return &http.Response{Body: body, StatusCode: g.statusCode}, nil
}

func TestSpikeClient_GetCompetitions(t *testing.T) {
	mockLeagueID := uint(124)
	mockStatus := "mockStatus"

	type args struct {
		leagueID uint
		status   string
	}
	tests := []struct {
		name    string
		c       *SpikeClient
		args    args
		want    []byte
		wantErr bool
	}{
		{"Success", NewSpikeClient(&mockConfig, &MockDoer{nil, 200}), args{mockLeagueID, mockStatus}, []byte("?status=" + mockStatus), false},
		{"Failure", NewSpikeClient(&mockConfig, &MockDoer{&errMock, 200}), args{}, nil, true},
		{"Invalid Status Code", NewSpikeClient(&mockConfig, &MockDoer{nil, 0}), args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCompetitions(tt.args.leagueID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpikeClient.GetCompetitions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpikeClient.GetCompetitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpikeClient_makeGetRequest(t *testing.T) {
	mockURL := "mockurl"

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		c       *SpikeClient
		args    args
		want    []byte
		wantErr bool
	}{
		{"Success", NewSpikeClient(&mockConfig, &MockDoer{nil, 200}), args{mockURL}, []byte(mockURL), false},
		{"Failure", NewSpikeClient(&mockConfig, &MockDoer{&errMock, 200}), args{mockURL}, nil, true},
		{"Invalid Status Code", NewSpikeClient(&mockConfig, &MockDoer{nil, 0}), args{mockURL}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.makeGetRequest(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpikeClient.makeGetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpikeClient.makeGetRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpikeClient_GetContests(t *testing.T) {
	mockCompetitionID := uint(456)
	mockStatus := uint(27)

	type args struct {
		competitionID uint
		status        uint
	}
	tests := []struct {
		name    string
		c       *SpikeClient
		args    args
		want    []byte
		wantErr bool
	}{
		{"Success", NewSpikeClient(&mockConfig, &MockDoer{nil, 200}), args{mockCompetitionID, mockStatus}, []byte("?status=27&started=1&finished=1"), false},
		{"Failure", NewSpikeClient(&mockConfig, &MockDoer{&errMock, 200}), args{mockCompetitionID, mockStatus}, nil, true},
		{"Invalid Status Code", NewSpikeClient(&mockConfig, &MockDoer{nil, 400}), args{mockCompetitionID, mockStatus}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetContests(tt.args.competitionID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpikeClient.GetContests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpikeClient.GetContests() = %v, want %v", got, tt.want)
			}
		})
	}
}
