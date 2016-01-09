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
	err := r.Retriever("GET", "me", &output)
	return output, err
}

func (r *Reddit) MeFriends() (UserList, error) {
	var output UserList
	err := r.Retriever("GET", "me/friends", &output)
	return output, err
}

func (r *Reddit) MeKarma() (KarmaList, error) {
	var output KarmaList
	err := r.Retriever("GET", "me/karma", &output)
	return output, err
}

func (r *Reddit) MePrefs() (Preferences, error) {
	var output Preferences
	err := r.Retriever("GET", "me/prefs", &output)
	return output, err
}

func (r *Reddit) MeTrophies() (TrophyList, error) {
	var output TrophyList
	err := r.Retriever("GET", "me/trophies", &output)
	return output, err
}
