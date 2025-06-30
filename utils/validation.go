package utils

import (
	"errors"
	"net/url"
)

func ValidateUrl(input string) (*url.URL, error) {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, errors.New("Invalid URL scheme")
	}

	return parsedURL, nil
}
