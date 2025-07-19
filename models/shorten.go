package models

import (
	"bufio"
	"errors"
	"example.com/url-shortner/db"
	"example.com/url-shortner/utils"
	"mime/multipart"
	"strings"
)

var UrlStore = make(map[string]string)

type Link struct {
	ID     int64  `json:"id"`
	Url    string `json:"url" binding:"required"`
	Code   string `json:"code"`
	UserID int64  `json:"user_id"`
}

type ShortenResult struct {
	URL  string `json:"url"`
	Code string `json:"code,omitempty"`
	Err  string `json:"err,omitempty"`
}

func (link *Link) ShortenUrl() (map[string]string, error) {
	err := utils.ValidateUrl(link.Url)
	if err != nil {
		return nil, err
	}

	code, err := utils.GenerateShortCode(6)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO links(code,url,user_id) VALUES ($1,$2,$3)`
	_, err = db.DB.Exec(query, code, link.Url, link.UserID)
	if err != nil {
		return nil, err
	}

	return map[string]string{"code": code}, nil
}

func (link Link) BulkUploadUrls(file multipart.File) ([]ShortenResult, error) {
	var results []ShortenResult

	scanner := bufio.NewScanner(file)
	stmt, err := db.DB.Prepare(`INSERT INTO links(code, url, user_id) VALUES ($1, $2, $3)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		err := utils.ValidateUrl(line)
		if err != nil {
			results = append(results, ShortenResult{
				URL: line,
				Err: "Invalid Url",
			})
			continue
		}

		var code string
		for {
			code, err = utils.GenerateShortCode(6)
			if err != nil {
				return nil, err
			}

			var exists bool
			err = db.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM links WHERE code = $1)`, code).Scan(&exists)
			if err != nil {
				return nil, err
			}

			if !exists {
				break
			}
		}

		_, err = stmt.Exec(code, line, link.UserID)
		if err != nil {
			return nil, err
		}

		results = append(results, ShortenResult{URL: line, Code: code})
	}

	if err := scanner.Err(); err != nil {
		results = append(results, ShortenResult{
			Err: "Scanner error: " + err.Error(),
		})
	}

	return results, nil
}

func (link *Link) GetOriginalUrl() error {
	query := "SELECT url FROM links WHERE code=$1;"
	row := db.DB.QueryRow(query, link.Code)
	err := row.Scan(&link.Url)
	if err != nil {
		return errors.New("url not found")
	}

	return nil
}
