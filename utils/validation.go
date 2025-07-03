package utils

import (
	"errors"
	"net/url"
)

func ValidateUrl(input string) error {
	parsedUrl, err := url.Parse(input)
	if err != nil {
		return err
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return errors.New("Invalid URL")
	}

	return nil
}
