package wrapper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// AccountLeaderboard gets the current accounts leaderboard position in the tickets leaderboard, this requires an authorization token to be set and **WILL** error if one is not provided
func (s *Session) AccountLeaderboard() (*AccountLeaderboard, error) {
	if s.Auth == "" {
		return nil, errors.New("no auth token set")
	}
	c := http.Client{}
	req, err := http.NewRequest("GET", AccountLeaderboardURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("cookie", "token="+s.Auth)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("http request failed")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &AccountLeaderboard{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetLeaderboard fetches the data for the current users in the tickets leaderboard, this requires no authorization and therefor doesn't use any cookie headers
func (s *Session) GetLeaderboard() (*TicketsLeaderboardResult, error) {
	c := http.Client{}
	req, err := http.NewRequest("GET", TicketsLeaderboardURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("http request failed")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &TicketsLeaderboard{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if !r.Success {
		return nil, errors.New("success was false")
	}
	return &r.Result, nil
}
