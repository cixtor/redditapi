package main

import (
	"github.com/cixtor/redditapi"
	"log"
)

func main() {
	cli := redditapi.New()

	if _, err := cli.Credentials(); err != nil {
		log.Println(err)
		log.Println("Info: https://www.reddit.com/prefs/apps")
		log.Println("Docs: https://www.reddit.com/wiki/api")
		return
	}

	log.Printf("%#v\n", cli)
}
