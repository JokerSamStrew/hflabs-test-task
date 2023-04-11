package main

import (
	"log"
	"os"
)

const docId = "1sY0nvgfcSFW7Y8j58gDhXNkLRtRHxQ_DgNQKwX8Lgqw"

func main() {
	// codes, err := getCurrentTable()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve response codes info %v", err)
	// }

	// fmt.Println(codes)

	codes, err := os.ReadFile("errorcodes.html") // just pass the file name
	if err != nil {
		panic(err)
	}

	// for i, item := range parseResponseTable(string(codes)) {
	// 	fmt.Printf("%v %v\t\t%v\n", i, item.ResponseCode, item.Description)
	// }

	srv, err := getdocaccess(docId, "client_secret.json")
	if err != nil {
		log.Fatalf("Unable gain access to document: %v", err)
	}

	// response, err := createTable(srv, docId, []TableRow{
	// 	{"Left1", "Right2"},
	// 	{"Left12", "Rightfasd"},
	// 	{"Left134", "Right1"},
	// 	{"Left11", "Right1"},
	// 	{"Left1", "Right1"},
	// 	{"Left1", "Right111"},
	// 	{"Left1fdasfdas", "Right1fdafasd"},
	// 	{"Left1", "Right1"},
	// })

	response, err := createTable(srv, docId, parseResponseTable(string(codes)))
	if err != nil {
		log.Fatalf("table creation fail: %v", err)
	}

	log.Println(response)
}
