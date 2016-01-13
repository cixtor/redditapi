package redditapi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const OAUTH_API = "https://oauth.reddit.com"
const BASIC_API = "https://www.reddit.com/api/v1/"
const USER_AGENT = "RedditAPI/0.1 by cixtor (+github.com/cixtor/redditapi)"

type Reddit struct {
	Client    string
	Secret    string
	Username  string
	Password  string
	UserAgent string
}

type Session struct {
	Code        int    `json:"error"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	CreatedAt   int    `json:"created_at"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func New() *Reddit {
	return &Reddit{UserAgent: USER_AGENT}
}

func (r *Reddit) Configure(client string, secret string) {
	r.Client = client
	r.Secret = secret
}

func (r *Reddit) Authorize(username string, password string) {
	r.Username = username
	r.Password = password
}

func (r *Reddit) SetUserAgent(unique string) {
	r.UserAgent = unique
}

func (r *Reddit) Token() (Session, error) {
	token, err := r.TokenFromFile()

	if err != nil {
		return r.TokenFromAPI()
	}

	return token, nil
}

func (r *Reddit) TokenFilepath() string {
	return "/tmp/redditapi.json"
}

func (r *Reddit) HasExpired(token Session) bool {
	return (int(time.Now().Unix()) - token.CreatedAt) >= token.ExpiresIn
}

func (r *Reddit) TokenFromFile() (Session, error) {
	var token Session

	filepath := r.TokenFilepath()
	file, err := os.Open(filepath)

	if err != nil {
		return token, err
	}

	if err := json.NewDecoder(file).Decode(&token); err != nil {
		return token, err
	}

	if r.HasExpired(token) {
		return token, errors.New("Token has expired")
	}

	return token, nil
}

func (r *Reddit) TokenFromAPI() (Session, error) {
	var token Session

	params := url.Values{}
	client := &http.Client{}

	params.Add("grant_type", "password")
	params.Add("username", r.Username)
	params.Add("password", r.Password)

	req, err := http.NewRequest("POST",
		BASIC_API+"access_token",
		bytes.NewBufferString(params.Encode()))

	if err != nil {
		return token, err
	}

	req.SetBasicAuth(r.Client, r.Secret)

	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,*/*;q=0.8")
	req.Header.Set("User-Agent", r.UserAgent)

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

	token.CreatedAt = int(time.Now().Unix())
	file, _ := os.Create(r.TokenFilepath())
	result, _ := json.Marshal(token)
	file.Write(result)

	return token, nil
}

func (r *Reddit) Request(method string, action string, params []string) (io.Reader, error) {
	token, err := r.Token()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, OAUTH_API+action, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,*/*;q=0.8")
	req.Header.Set("Authorization", token.TokenType+"\x20"+token.AccessToken)
	req.Header.Set("User-Agent", r.UserAgent)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	(&buf).ReadFrom(resp.Body)

	return &buf, nil
}

func (r *Reddit) RequestString(method string, action string, params []string) []byte {
	var chunk []byte
	var output []byte

	stream, err := r.Request(method, action, nil)

	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(stream)

	for scanner.Scan() {
		chunk = scanner.Bytes()
		output = append(output, chunk...)
	}

	return output
}

func (r *Reddit) RequestJson(method string, action string, params []string, output interface{}) error {
	stream, err := r.Request(method, action, nil)

	if err != nil {
		return err
	}

	return json.NewDecoder(stream).Decode(&output)
}

func (r *Reddit) Get(action string, params []string) []byte {
	return r.RequestString("GET", action, params)
}

func (r *Reddit) Post(action string, params []string) []byte {
	return r.RequestString("POST", action, params)
}

func (r *Reddit) GetJson(action string, params []string, output interface{}) error {
	return r.RequestJson("GET", action, params, &output)
}

func (r *Reddit) PostJson(action string, params []string, output interface{}) error {
	return r.RequestJson("POST", action, params, &output)
}
