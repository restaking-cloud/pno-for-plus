package common

import (
	"net/url"
	"strings"
)

func CreateUrl(urlString string) (*url.URL, error) {
	if urlString == "" {
		return nil, nil
	}
	if !strings.HasPrefix(urlString, "http") {
		urlString = "http://" + urlString
	}

	return url.ParseRequestURI(urlString)
}