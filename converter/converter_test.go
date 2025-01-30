package converter_test

import (
	_ "embed"
	"testing"

	"github.com/cpustejovsky/quotes/converter"
)

//go:embed quotes.test.txt
var quotes string

func TestConvertQuotes(t *testing.T) {
	converted, err := converter.ConvertQuotes(quotes)
	if err != nil {
		t.Fatal(err)
	}
	expect := 10
	got := len(converted)
	if got != expect {
		t.Fatalf("got %d, expected %d", got, expect)
	}

}
