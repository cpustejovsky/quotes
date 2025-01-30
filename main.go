package main

import (
	"log"
	"math/rand/v2"
	"os/exec"

	"github.com/cpustejovsky/quotes/converter"
)

func main() {
	quotes, err := converter.ConvertQuotes("/home/cpustejovsky/development/go/quotes/news-feed-eradicator.txt")
	if err != nil {
		log.Fatal(err)
	}
	index := rand.IntN(len(quotes) - 1)
	cmd := exec.Command("notify-send", quotes[index])
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
