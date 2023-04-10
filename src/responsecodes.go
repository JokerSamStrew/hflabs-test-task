package main

import (
	"encoding/json"
	"fmt"
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

func parseResponseTable(responseTable string) string {
	tkn := html.NewTokenizer(strings.NewReader(responseTable))

	for {
		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return "end"
		case tt == html.StartTagToken:
			t := tkn.Token()
			switch {
			case t.Data == "tr":
				fmt.Println()
			case t.Data == "td" || t.Data == "th":
				fmt.Printf("\t")
			}

		case tt == html.TextToken:
			t := tkn.Token()
			text := strings.TrimSpace(t.Data)
			if text == "" {
				continue
			}

			fmt.Printf("%v", text)
		}
	}
}
