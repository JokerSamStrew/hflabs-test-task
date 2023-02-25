package main

import (
	"encoding/json"
	"io"
	"net/http"
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
