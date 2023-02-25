package main

import (
	"fmt"
	"log"
)

const docId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

func main() {
	doc, err := getdocaccess(docId, "client_secret.json")
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	fmt.Println(doc.Title)

	codes, err := getCurrentTable()
	if err != nil {
		log.Fatalf("Unable to retrieve response codes info %v", err)
	}

	fmt.Println(codes)
}
