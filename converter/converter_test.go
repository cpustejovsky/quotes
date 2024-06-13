package converter_test

import (
	"github.com/cpustejovsky/quotes/converter"
	"testing"
)

func TestConvertQuotes(t *testing.T) {
	converted, err := converter.ConvertQuotes("/home/cpustejovsky/go/src/quotes/news-feed-eradicator.md")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(converted)
}
