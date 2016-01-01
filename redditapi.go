package redditapi

import (
	"errors"
	"os"
)

type Reddit struct {
	Client   string
	Secret   string
	Username string
	Password string
}

func New() *Reddit {
	return &Reddit{}
}

func (r *Reddit) Credentials() (bool, error) {
	r.Client = os.Getenv("REDDIT_CLIENT")
	r.Secret = os.Getenv("REDDIT_SECRET")
	r.Username = os.Getenv("REDDIT_USERNAME")
	r.Password = os.Getenv("REDDIT_PASSWORD")

	if r.Client == "" {
		return false, errors.New("Missing REDDIT_CLIENT environment variable")
	}

	if r.Secret == "" {
		return false, errors.New("Missing REDDIT_SECRET environment variable")
	}

	if r.Username == "" {
		return false, errors.New("Missing REDDIT_USERNAME environment variable")
	}

	if r.Password == "" {
		return false, errors.New("Missing REDDIT_PASSWORD environment variable")
	}

	return true, nil
}
