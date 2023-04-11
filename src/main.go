package main

import (
	"log"
)

const docId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

func main() {
	codes, err := getCurrentTable()
	if err != nil {
		log.Fatalf("Unable to retrieve response codes info %v", err)
	}

	srv, err := getdocaccess(docId, "client_secret.json")
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	response, err := createTable(srv, docId, parseResponseTable(string(codes)))
	if err != nil {
		log.Fatalf("Table creation fail: %v", err)
	}

	log.Println(response)
}
