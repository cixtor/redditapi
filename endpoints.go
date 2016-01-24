package redditapi

import "io"

type Endpoints interface {
	Me() (Account, error)
	MeFriends() (UserList, error)
	MeKarma() (KarmaList, error)
	MePrefs() (Preferences, error)
	MeTrophies() (TrophyList, error)
	NeedsCaptcha() (bool, error)
	NewCaptcha() []byte
	CaptchaIden(string) (io.Reader, error)
	Comment(string, string) (Comment, error)
	Delete(string) error
	Edit(string, string) (Comment, error)
	Hide(string) error
	Unhide(string) error
	Info(string, string) (Info, error)
	Lock(string) error
	Unlock(string) error
	MarkNSFW(string) error
	UnmarkNSFW(string) error
}

func (r *Reddit) Me() (Account, error) {
	var output Account
	err := r.GetJson("/api/v1/me", nil, &output)
	return output, err
}

func (r *Reddit) MeFriends() (UserList, error) {
	var output UserList
	err := r.GetJson("/api/v1/me/friends", nil, &output)
	return output, err
}

func (r *Reddit) MeKarma() (KarmaList, error) {
	var output KarmaList
	err := r.GetJson("/api/v1/me/karma", nil, &output)
	return output, err
}

func (r *Reddit) MePrefs() (Preferences, error) {
	var output Preferences
	err := r.GetJson("/api/v1/me/prefs", nil, &output)
	return output, err
}

func (r *Reddit) MeTrophies() (TrophyList, error) {
	var output TrophyList
	err := r.GetJson("/api/v1/me/trophies", nil, &output)
	return output, err
}

func (r *Reddit) NeedsCaptcha() (bool, error) {
	var output bool
	err := r.GetJson("/api/needs_captcha", nil, &output)
	return output, err
}

func (r *Reddit) NewCaptcha() []byte {
	return r.Get("/api/new_captcha", nil)
}

func (r *Reddit) CaptchaIden(iden string) (io.Reader, error) {
	return r.Request("GET", "/captcha/"+iden, nil)
}

func (r *Reddit) Comment(parent string, text string) (Comment, error) {
	var output Comment
	err := r.PostJson("/api/comment", map[string]string{
		"api_type": "json",
		"parent":   parent,
		"text":     text,
	}, &output)
	return output, err
}

func (r *Reddit) Delete(thing string) error {
	return r.PostJson("/api/del", map[string]string{"id": thing}, nil)
}

func (r *Reddit) Edit(thing string, text string) (Comment, error) {
	var output Comment
	err := r.PostJson("/api/editusertext", map[string]string{
		"api_type": "json",
		"thing_id": thing,
		"text":     text,
	}, &output)
	return output, err
}

func (r *Reddit) Hide(thing string) error {
	return r.PostJson("/api/hide", map[string]string{"id": thing}, nil)
}

func (r *Reddit) Unhide(thing string) error {
	return r.PostJson("/api/unhide", map[string]string{"id": thing}, nil)
}

func (r *Reddit) Info(thing string, url string) (Info, error) {
	var output Info
	err := r.GetJson("/api/info", map[string]string{"id": thing, "url": url}, &output)
	return output, err
}

func (r *Reddit) Lock(thing string) error {
	return r.PostJson("/api/lock", map[string]string{"id": thing}, nil)
}

func (r *Reddit) Unlock(thing string) error {
	return r.PostJson("/api/unlock", map[string]string{"id": thing}, nil)
}

func (r *Reddit) MarkNSFW(thing string) error {
	return r.PostJson("/api/marknsfw", map[string]string{"id": thing}, nil)
}

func (r *Reddit) UnmarkNSFW(thing string) error {
	return r.PostJson("/api/unmarknsfw", map[string]string{"id": thing}, nil)
}
