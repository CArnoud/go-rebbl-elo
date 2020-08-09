package api

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/CArnoud/go-rebbl-elo/config"
)

// Doer interface to make HTTP GET requests.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// SpikeClient client to access Spike endpoints
type SpikeClient struct {
	config *config.Config
	doer   Doer
}

func (c *SpikeClient) makeGetRequest(url string) ([]byte, error) {
	log.Println("Making HTTP request to " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("API_KEY", c.config.SpikeAPIKey)
	resp, err := c.doer.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("HTTP error: " + resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetCompetitions returns a list of competition IDs for a league.
func (c *SpikeClient) GetCompetitions(leagueID uint, status string) ([]byte, error) {
	url := c.config.SpikeAPIHost + c.config.SpikeCompetitionsPath
	url = strings.Replace(url, "<id>", strconv.FormatUint(uint64(leagueID), 10), 1)
	url = url + "?status=" + status
	body, err := c.makeGetRequest(url)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetContests returns a list of matches in a specific competition.
func (c *SpikeClient) GetContests(competitionID uint, status uint) ([]byte, error) {
	url := c.config.SpikeAPIHost + c.config.SpikeContestsPath
	url = strings.Replace(url, "<id>", strconv.FormatUint(uint64(competitionID), 10), 1)
	url = url + "?status=" + strconv.FormatUint(uint64(status), 10)
	url = url + "&started=1&finished=1"
	body, err := c.makeGetRequest(url)
	if err != nil {
		return nil, err
	}

	return body, err
}

// NewSpikeClient creates an instance of SpikeClient.
func NewSpikeClient(config *config.Config, doer Doer) *SpikeClient {
	return &SpikeClient{config, doer}
}
