package main

import (
	_ "embed"
	"log"
	"math/rand/v2"
	"os/exec"

	"github.com/cpustejovsky/quotes/converter"
)

//go:embed quotes.txt
var quotes string

func main() {
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
