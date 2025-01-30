package converter

type Quote struct {
	Text        string
	Attribution string
}

func ConvertQuotes(dat string) ([]string, error) {
	var quotes []Quote
	var q Quote
	var attributionNeed bool
	for _, char := range string(dat) {
		if char == '~' {
			attributionNeed = true
			continue
		}
		if attributionNeed {
			if char == '\n' {
				quotes = append(quotes, q)
				q = Quote{}
				attributionNeed = false
			} else {
				q.Attribution += string(char)
			}
		} else {
			q.Text += string(char)
		}
	}
	var converted []string
	for _, quote := range quotes {
		converted = append(converted, "\""+quote.Text+"\"\n-"+quote.Attribution+"\n")
	}
	return converted, nil
}
