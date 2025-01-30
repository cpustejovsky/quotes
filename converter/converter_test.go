package converter_test

import (
	"testing"

	"github.com/cpustejovsky/quotes/converter"
)

func TestConvertQuotes(t *testing.T) {
	converted, err := converter.ConvertQuotes("/home/cpustejovsky/development/go/quotes/news-feed-eradicator.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(converted)
}
