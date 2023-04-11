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

	// codes, err := os.ReadFile("errorcodes.html") // just pass the file name
	// if err != nil {
	// 	panic(err)
	// }

	// for i, item := range parseResponseTable(string(codes)) {
	// 	fmt.Printf("%v %v\t\t%v\n", i, item.ResponseCode, item.Description)
	// }

	srv, err := getdocaccess(docId, "client_secret.json")
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	response, err := createTable(srv, docId, []TableRow{
		{"Left1", "Right1"},
		{"Left1", "Right1"},
		{"Left1", "Right1"},
		{"Left1", "Right1"},
		{"Left1", "Right1"},
		{"Left1", "Right1"},
	})
	if err != nil {
		log.Fatalf("table creation fail: %v", err)
	}

	log.Println(response)
}
