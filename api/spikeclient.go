package api

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/CArnoud/go-rebbl-elo/config"
)

// Getter interface to make HTTP GET requests.
type Getter interface {
	Get(string) (*http.Response, error)
}

// SpikeClient client to access Spike endpoints
type SpikeClient struct {
	config *config.Config
	getter Getter
}

func (c *SpikeClient) makeCompetitionsURL(leagueID uint) string {
	url := c.config.SpikeAPIHost + c.config.SpikeCompetitionsPath
	url = url + "?league_id=" + strconv.FormatUint(uint64(leagueID), 10)
	return url
}

// GetCompetitions returns a list of competition IDs for a league.
func (c *SpikeClient) GetCompetitions(leagueID uint) ([]byte, error) {
	resp, err := c.getter.Get(c.makeCompetitionsURL(leagueID))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// NewSpikeClient creates an instance of SpikeClient.
func NewSpikeClient(config *config.Config, getter Getter) *SpikeClient {
	return &SpikeClient{config, getter}
}
