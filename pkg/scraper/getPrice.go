package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetPrice(url string) (string, error) {
	// Check if URL is valid amazon URL
	urlIsValid := CheckURL(url)
	if !urlIsValid {
		return "", fmt.Errorf("invalid amazon url")
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set a User-Agent to mimic a browser (this helps avoid some anti-scraping measures)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	// Function to recursively search for price elements
	// var unit string
	var dollars string
	var cents string
	var dollarsFound bool = false
	var centsFound bool = false

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if dollarsFound && centsFound {
					return
				}

				if !dollarsFound && a.Key == "class" && strings.Contains(a.Val, "a-price-whole") {
					dollars = n.FirstChild.Data
					dollarsFound = true
				}
				if !centsFound && a.Key == "class" && strings.Contains(a.Val, "a-price-fraction") {
					cents = n.FirstChild.Data
					centsFound = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if cents == "" || dollars == "" {
		return "", fmt.Errorf("price not found on the page")
	}

	return "$" + dollars + "." + cents, nil
}
