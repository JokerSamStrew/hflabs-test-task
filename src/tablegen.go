package main

import (
	"fmt"

	"google.golang.org/api/docs/v1"
)

func createTable(srv *docs.Service) (*docs.BatchUpdateDocumentResponse, error) {
	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from document: %v", err)
	}

	var maxEndIndex int64 = 0
	for _, se := range doc.Body.Content {
		if maxEndIndex < se.EndIndex {
			maxEndIndex = se.EndIndex
		}
	}

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
	request := docs.BatchUpdateDocumentRequest{Requests: requests}

	response, err := srv.Documents.BatchUpdate(docId, &request).Do()
	if err != nil {
		return nil, fmt.Errorf("BatchUpdate failed: %v", err)
	}

	return response, nil
}
