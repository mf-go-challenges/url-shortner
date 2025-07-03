package models

import (
	"bufio"
	"example.com/url-shortner/db"
	"example.com/url-shortner/utils"
	"mime/multipart"
	"strings"
)

var UrlStore = make(map[string]string)

type ShortUrl struct {
	Url string `json:"url" binding:"required"`
}

type ShortenResult struct {
	URL  string `json:"url"`
	Code string `json:"code"`
}

func (url ShortUrl) ShortenUrl() (map[string]string, error) {
	err := utils.ValidateUrl(url.Url)
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

func BulkUploadUrls(file multipart.File) ([]ShortenResult, error) {
	var results []ShortenResult
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		err := utils.ValidateUrl(line)
		if err != nil {
			continue
		}
		code, err := utils.GenerateShortCode(6)
		if err != nil {
			return nil, err
		}
		stmt, err := db.DB.Prepare(`INSERT INTO links(code, url, created_at) VALUES (?, ?, CURRENT_TIMESTAMP)`)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(code, line)
		if err != nil {
			return nil, err
		}

		results = append(results, ShortenResult{URL: line, Code: code})
	}

	return results, nil
}
