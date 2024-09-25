package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetItemName(url string) (string, error) {
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

	// Function to recursively search for the item name
	var itemName string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "productTitle" {
					itemName = strings.TrimSpace(n.FirstChild.Data)
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if itemName == "" {
		return "", fmt.Errorf("item name not found on the page")
	}

	return itemName, nil
}
