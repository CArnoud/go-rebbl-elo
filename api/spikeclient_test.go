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

type MockGetter struct{}

func (g *MockGetter) Get(url string) (*http.Response, error) {
	if url == "?league_id=666" {
		return nil, errors.New("mockerror")
	}

	body := ioutil.NopCloser(bytes.NewBufferString(url))
	return &http.Response{Body: body}, nil
}

func TestSpikeClient_GetCompetitions(t *testing.T) {
	mockLeagueID := uint(124)
	errorLeagueID := uint(666)

	mockConfig := config.Config{}

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
		{"Success", NewSpikeClient(&mockConfig, &MockGetter{}), args{mockLeagueID}, []byte("?league_id=124"), false},
		{"Failure", NewSpikeClient(&mockConfig, &MockGetter{}), args{errorLeagueID}, nil, true},
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
