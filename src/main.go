package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	docs "google.golang.org/api/docs/v1"
)

const ResponseCodesUrl = "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"
const DocumentId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

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

func parseHtmlProcedure() {
	respCodesHtml, err := retrieveResponseCodesHtmlPage()
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	fmt.Println(respCodesHtml)
}

func main() {
	ctx := context.Background()
	fmt.Printf("ctx %v\n", ctx)
	docsService, err := docs.NewService(ctx)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	fmt.Printf("docsService %v\n", docsService)
	insertTableRequest := docs.InsertTableRequest{Columns: 5, EndOfSegmentLocation: &docs.EndOfSegmentLocation{}, Location: &docs.Location{}}

	fmt.Printf("insertTableRequest %v\n", insertTableRequest)
	requests := []*docs.Request{{InsertTable: &insertTableRequest}}

	batchUpdate := docs.BatchUpdateDocumentRequest{Requests: requests, WriteControl: &docs.WriteControl{}}
	fmt.Printf("batchUpdate %v\n", batchUpdate)
	response, err := docsService.Documents.BatchUpdate(DocumentId, &batchUpdate).Context(ctx).Do()
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	fmt.Printf("response %v\n", response)
}
