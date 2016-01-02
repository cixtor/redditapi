package redditapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Reddit struct {
	Client    string
	Secret    string
	Username  string
	Password  string
	UserAgent string
}

type Session struct {
	Code        int    `json:"error,omitempty"`
	Message     string `json:"message,omitempty"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	CreatedAt   int    `json:"created_at"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func New() *Reddit {
	return &Reddit{}
}

func (r *Reddit) Configure(client string, secret string) {
	r.Client = client
	r.Secret = secret
}

func (r *Reddit) Authorize(username string, password string) {
	r.Username = username
	r.Password = password
}

func (r *Reddit) TokenFromAPI() (Session, error) {
	var token Session

	params := url.Values{}
	client := &http.Client{}

	params.Add("grant_type", "password")
	params.Add("username", r.Username)
	params.Add("password", r.Password)

	req, err := http.NewRequest("POST",
		"https://www.reddit.com/api/v1/access_token",
		bytes.NewBufferString(params.Encode()))

	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,*/*;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", r.UserAgent)

	if err != nil {
		return token, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return token, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return token, err
	}

	if token.Code != 0 {
		return Session{}, errors.New(fmt.Sprintf("(%d) %s", token.Code, token.Message))
	}

	return token, nil
}
