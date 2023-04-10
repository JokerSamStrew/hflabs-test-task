package main

import (
	"fmt"
	"os"
)

const docId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

func main() {
	// codes, err := getCurrentTable()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve response codes info %v", err)
	// }

	// fmt.Println(codes)

	// doc, err := getdocaccess(docId, "client_secret.json")
	// if err != nil {
	// 	log.Fatalf("Unable gain access to document: %v", err)
	// }

	// fmt.Println(doc.Title)

	b, err := os.ReadFile("errorcodes.html") // just pass the file name
	if err != nil {
		panic(err)
	}

	fmt.Println(parseResponseTable(string(b)))
}
