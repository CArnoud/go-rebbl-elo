package api

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// Getter interface to make HTTP GET requests.
type Getter interface {
	Get(string) (*http.Response, error)
}

// SpikeClient client to access Spike endpoints
type SpikeClient struct {
	getter Getter
}

// GetCompetitions returns a list of competition IDs for a league.
func (c *SpikeClient) GetCompetitions(leagueID uint) ([]byte, error) {
	resp, err := c.getter.Get("" + strconv.FormatUint(uint64(leagueID), 10))
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
func NewSpikeClient(getter Getter) *SpikeClient {
	return &SpikeClient{getter}
}
