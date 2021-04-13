package wrapper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// MakeRequest builds a new request, it's to chop down on reused code
func (s *Session) MakeRequest(auth bool, method, URL string, body *strings.Reader) (*http.Request, error) {
	if auth {
		if s.Auth == "" {
			return nil, errors.New("no auth token set")
		}
		req, err := http.NewRequest(method, URL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("cookie", "token="+s.Auth)
		return req, nil
	}
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetBody does all the misc checking and returns the byte body of an http request
func (s *Session) GetBody(req *http.Request) ([]byte, error) {
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, errors.New("http request failed with code " + resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

// AllInOneHTTP is just a wrap of MakeRequest and GetBody to chop down on reused code while letting me keep control in the future
func (s *Session) AllInOneHTTP(auth bool, Method, URL string, body *strings.Reader) ([]byte, error) {
	req, err := s.MakeRequest(auth, Method, URL, body)
	if err != nil {
		return []byte{}, err
	}
	b, err := s.GetBody(req)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

// AccountLeaderboard gets the current accounts leaderboard position in the tickets leaderboard, this requires an authorization token to be set and **WILL** error if one is not provided
func (s *Session) AccountLeaderboard() (*AccountLeaderboard, error) {
	b, err := s.AllInOneHTTP(true, "GET", AccountLeaderboardURL, nil)
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
	b, err := s.AllInOneHTTP(true, "GET", AccountLeaderboardURL, nil)
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

// AccountEarnings is the func to get accounts earnings, it returns the amount of money put in and the amount of money won so you can calculate a profit amount. You **need** auth for this func, if auth isn't set it will give back an error
func (s *Session) AccountEarnings() (*TotalWageredResult, error) {
	req, err := s.MakeRequest(true, "GET", AccountEarningsURL, nil)
	if err != nil {
		return nil, err
	}
	b, err := s.GetBody(req)
	if err != nil {
		return nil, err
	}
	r := &TotalWagered{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if !r.Success {
		return nil, errors.New("success was false")
	}
	return &r.Result, nil
}
