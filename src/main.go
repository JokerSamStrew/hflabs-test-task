package main

import (
	"log"
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

	response, err := createTable(srv)
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	log.Println(response)
}
