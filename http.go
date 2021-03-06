package wrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// MakeRequest builds a new request, it's to chop down on reused code
func (s *Session) MakeRequest(auth bool, method, URL string, body *strings.Reader) (*http.Request, error) {
	if auth {
		if s.Auth == "" {
			return nil, errors.New("no auth token set")
		}
		var req *http.Request
		var err error
		if body == nil {
			req, err = http.NewRequest(method, URL, nil)
		} else {
			req, err = http.NewRequest(method, URL, body)
		}
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

// AccountJSONRegex is the regex to pull the json out of the profile HTML
var AccountJSONRegex = regexp.MustCompile(`window\.userData={.+}`)
var getJSONCorrect = regexp.MustCompile(`[A-z]+:`)

func formatJSON(s string) string {
	final := s
	matches := getJSONCorrect.FindAllString(s, -1)
	for _, match := range matches {
		if !strings.Contains(match, "http") {
			final = strings.Replace(final, match, fmt.Sprintf("\"%s\":", strings.ReplaceAll(match, ":", "")), 1)
		}
	}
	return final
}

// GetAccountInfo gets the general information of an account
func (s *Session) GetAccountInfo() (*AccountInfo, error) {
	req, err := s.MakeRequest(true, "GET", AccountProfileURL, nil)
	if err != nil {
		return nil, err
	}
	b, err := s.GetBody(req)
	if err != nil {
		return nil, err
	}
	matches := AccountJSONRegex.FindAllString(string(b), -1)
	if len(matches) < 1 {
		return nil, err
	}
	match := strings.ReplaceAll(matches[0], "window.userData=", "")
	r := &AccountInfo{}
	match = formatJSON(match)
	err = json.Unmarshal([]byte(match), &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// ClaimFaucet attempts to claim the free 3 cent faucet
// CaptchaToken is an hcaptcha token, you'd need to use 2captcha or a similar service to get this CaptchaToken
// Response is a *FaucetResponse, followed by an error
func (s *Session) ClaimFaucet(CaptchaToken string) (*FaucetResponse, error) {
	req, err := http.NewRequest("POST", FaucetClaimURL, strings.NewReader("response="+CaptchaToken))
	if err != nil {
		return nil, err
	}
	req.Header.Set("cookie", "token="+s.Auth)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &FaucetResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// CheckSerial returns the "Provably fair" response from a URL like https://rustchance.com/provably-fair/serial?number=5465806
// Serial with the example of that URL would be 5465806 aka the game number (not ID)
func (s *Session) CheckSerial(Serial string) (*ProvablyFair, error) {
	resp, err := s.AllInOneHTTP(false, "GET", ProvefairSerialURL+Serial, nil)
	if err != nil {
		return nil, err
	}
	r := &ProvablyFair{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// RedeemCode attempts to redeem a sponsored code on an account
// Code is the sponsored code, for example, "CHANCE"
// NOTE: I have no clue what a valid and successful response looks like...
func (s *Session) RedeemCode(Code string) (*RedeemCodeResponse, error) {
	req, err := http.NewRequest("POST", RedeemCodeURL, strings.NewReader("code="+Code))
	if err != nil {
		return nil, err
	}
	req.Header.Set("cookie", "token="+s.Auth)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &RedeemCodeResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetCoinflipHistory retreeves the previous coinflip games
func (s *Session) GetCoinflipHistory() (*CoinflipHistory, error) {
	resp, err := s.AllInOneHTTP(false, "GET", HistoryAPIURL+"coinflip", nil)
	if err != nil {
		return nil, err
	}
	r := &CoinflipHistory{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetJackpotHistory retreeves the history of previous jackpot games
// Room is either low or high depending on which room you want
func (s *Session) GetJackpotHistory(Room string) (*JackpotHistory, error) {
	if Room != "low" && Room != "high" {
		return nil, errors.New("invalid room provided, must be high or low")
	}
	resp, err := s.AllInOneHTTP(false, "GET", HistoryAPIURL+"jackpot?room="+Room, nil)
	if err != nil {
		return nil, err
	}
	r := &JackpotHistory{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetCrashGame gets a coinflip game by ID
// ID is the ID of the coinflip game
func (s *Session) GetCrashGame(ID string) (*CrashGame, error) {
	resp, err := s.AllInOneHTTP(false, "GET", CrashGameURL+ID, nil)
	if err != nil {
		return nil, err
	}
	r := &CrashGame{}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
