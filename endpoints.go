package redditapi

type Endpoints interface {
	Me() (Account, error)
	MeFriends() (UserList, error)
	MeKarma() (KarmaList, error)
	MePrefs() (Preferences, error)
	MeTrophies() (TrophyList, error)
}

func (r *Reddit) Retriever(method string, action string, output interface{}) error {
	return r.Request(method, action, &output)
}

func (r *Reddit) Me() (Account, error) {
	var output Account
	err := r.Retriever("GET", "/api/v1/me", &output)
	return output, err
}

func (r *Reddit) MeFriends() (UserList, error) {
	var output UserList
	err := r.Retriever("GET", "/api/v1/me/friends", &output)
	return output, err
}

func (r *Reddit) MeKarma() (KarmaList, error) {
	var output KarmaList
	err := r.Retriever("GET", "/api/v1/me/karma", &output)
	return output, err
}

func (r *Reddit) MePrefs() (Preferences, error) {
	var output Preferences
	err := r.Retriever("GET", "/api/v1/me/prefs", &output)
	return output, err
}

func (r *Reddit) MeTrophies() (TrophyList, error) {
	var output TrophyList
	err := r.Retriever("GET", "/api/v1/me/trophies", &output)
	return output, err
}

func (r *Reddit) NeedsCaptcha() (bool, error) {
	var output bool
	err := r.Retriever("GET", "/api/needs_captcha", &output)
	return output, err
}
