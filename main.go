package main

import (
	_ "embed"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"

	"github.com/cpustejovsky/quotes/converter"
)

//go:embed quotes.txt
var quotes string

func main() {
	for _, arg := range os.Args {
		if arg == "prayer" {
			cmd := exec.Command("notify-send", "Be the helper of my soul, O god, for I walk in the midst of many snares. Deliver me from them and save me, for You are good, and You love mankind.")
			_, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}
	quotes, err := converter.ConvertQuotes(quotes)
	if err != nil {
		log.Fatal(err)
	}
	index := rand.IntN(len(quotes) - 1)
	cmd := exec.Command("notify-send", quotes[index])
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
