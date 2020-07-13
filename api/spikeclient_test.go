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
	errorLeagueID := uint(666)

	mockConfig := config.Config{}
	mockError := errors.New("mockerror")

	type args struct {
		leagueID uint
	}
	tests := []struct {
		name    string
		c       *SpikeClient
		args    args
		want    []byte
		wantErr bool
	}{
		{"Success", NewSpikeClient(&mockConfig, &MockDoer{nil, 200}), args{mockLeagueID}, []byte("?league_id=124"), false},
		{"Failure", NewSpikeClient(&mockConfig, &MockDoer{&mockError, 0}), args{errorLeagueID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCompetitions(tt.args.leagueID)
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
	mockConfig := config.Config{}
	mockURL := "mockurl"
	mockError := errors.New("mockerror")

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
		{"Failure", NewSpikeClient(&mockConfig, &MockDoer{&mockError, 200}), args{mockURL}, nil, true},
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
