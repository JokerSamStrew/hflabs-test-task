package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const ResponseTableUrl = "https://confluence.hflabs.ru/rest/api/content/1181220999?expand=body.storage"

func getCurrentTable() (string, error) {
	response, err := http.Get(ResponseTableUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var restData map[string]interface{}
	err = json.Unmarshal(body, &restData)
	if err != nil {
		return "", err
	}

	value := restData["body"].(map[string]interface{})["storage"].(map[string]interface{})["value"].(string)
	return value, err
}

type TableRow struct {
	ResponseCode string
	Description  string
}

func parseResponseTable(responseTable string) []TableRow {
	tkn := html.NewTokenizer(strings.NewReader(responseTable))
	rows := make([]TableRow, 0, 10)
	var row TableRow
	isResponseCode := false

	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return rows

		case tt == html.StartTagToken:
			t := tkn.Token()
			switch {
			case t.Data == "tr":
				if row.ResponseCode != "" || row.Description != "" {
					rows = append(rows, row)
					row = TableRow{}
				}
			case t.Data == "td" || t.Data == "th":
				isResponseCode = !isResponseCode
			case t.Data == "li":
				if isResponseCode {
					row.ResponseCode += "\n"
				} else {
					row.Description += "\n"
				}
			}

		case tt == html.TextToken:
			t := tkn.Token()
			text := strings.TrimSpace(t.Data)
			if text == "" {
				continue
			}

			if isResponseCode {
				row.ResponseCode += t.Data
			} else {
				row.Description += t.Data
			}
		}
	}

}
