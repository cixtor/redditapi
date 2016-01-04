package redditapi

type Endpoints interface {
	Me() Account
}

func (r *Reddit) Me() (Account, error) {
	var output Account

	err := r.Request("GET", "me", &output)

	if err != nil {
		return output, err
	}

	return output, nil
}
