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
	insertTableRequest := docs.InsertTableRequest{
		Location: &docs.Location{Index: 1},
		Rows:     6,
		Columns:  2,
	}
	request := docs.BatchUpdateDocumentRequest{}
	request.Requests = append(request.Requests, &docs.Request{InsertTable: &insertTableRequest})
	response, err := srv.Documents.BatchUpdate(docId, &request).Do()
	if err != nil {
		log.Fatalf("BatchUpdate failed: %v", err)
	}

	fmt.Println(response)
}
