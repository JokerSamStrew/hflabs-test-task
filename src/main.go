package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const ResponseCodesUrl = "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"

func retrieveResponseCodesHtmlPage() (string, error) {
	resp, httpErr := http.Get(ResponseCodesUrl)
	if httpErr != nil {
		return "", httpErr
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}

	return string(body), nil
}

func main() {
	respCodesHtml, err := retrieveResponseCodesHtmlPage()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(respCodesHtml)
}
