package main

import (
	"fmt"
	"log"

	"google.golang.org/api/docs/v1"
)

const docId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

func main() {
	// codes, err := getCurrentTable()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve response codes info %v", err)
	// }

	// fmt.Println(codes)

	// b, err := os.ReadFile("errorcodes.html") // just pass the file name
	// if err != nil {
	// 	panic(err)
	// }

	// for i, item := range parseResponseTable(codes) {
	// 	fmt.Printf("%v %v\t\t%v\n", i, item.ResponseCode, item.Description)
	// }

	srv, err := getdocaccess(docId, "client_secret.json")
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}

	fmt.Println(doc.Title)
	var maxEndIndex int64 = 0
	for _, se := range doc.Body.Content {
		if maxEndIndex < se.EndIndex {
			maxEndIndex = se.EndIndex
		}
	}
	fmt.Println(maxEndIndex)

	requests := []*docs.Request{}

	if maxEndIndex > 2 {
		requests = append(requests, &docs.Request{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{StartIndex: 1, EndIndex: maxEndIndex - 1},
			}})
	}

	insertTableRequest := docs.InsertTableRequest{
		Location: &docs.Location{Index: 1},
		Rows:     6,
		Columns:  2,
	}

	requests = append(requests, &docs.Request{InsertTable: &insertTableRequest})

	if len(requests) == 0 {
		log.Fatalf("No requests to send")
	}

	request := docs.BatchUpdateDocumentRequest{Requests: requests}

	response, err := srv.Documents.BatchUpdate(docId, &request).Do()
	if err != nil {
		log.Fatalf("BatchUpdate failed: %v", err)
	}

	fmt.Println(response)
}
