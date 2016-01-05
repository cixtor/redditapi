package main

import (
	"github.com/cixtor/redditapi"
	"log"
	"os"
)

func main() {
	cli := redditapi.New()

	cli.Configure(os.Getenv("REDDIT_CLIENT"), os.Getenv("REDDIT_SECRET"))
	cli.Authorize(os.Getenv("REDDIT_USERNAME"), os.Getenv("REDDIT_PASSWORD"))

	out, err := cli.Me()

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%#v\n", out)
}
