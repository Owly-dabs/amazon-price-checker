package scraper

import (
	"regexp"
)

func CheckURL(url string) bool {
	re := regexp.MustCompile(`https?:\/\/(www\.)?amazon\.([a-z.]{2,6})\/.*`)
	match := re.FindString(url)
	if match == "" {
		return false
	}
	return true
}
