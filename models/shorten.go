package models

import (
	"example.com/url-shortner/utils"
)

var UrlStore = make(map[string]string)

type ShortUrl struct {
	Url string `json:"url" binding:"required"`
}

func (url ShortUrl) ShortenUrl() (map[string]string, error) {
	_, err := utils.ValidateUrl(url.Url)
	if err != nil {
		return nil, err
	}

	code, err := utils.GenerateShortCode(6)
	if err != nil {
		return nil, err
	}

	UrlStore[code] = url.Url
	return map[string]string{"code": code}, nil
}
