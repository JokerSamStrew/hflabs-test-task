package main

import (
	"fmt"
	"unicode/utf8"

	"google.golang.org/api/docs/v1"
)

func clearDocRequest(srv *docs.Service, doc *docs.Document) *docs.Request {
	var maxEndIndex int64 = 0
	for _, se := range doc.Body.Content {
		if maxEndIndex < se.EndIndex {
			maxEndIndex = se.EndIndex
		}
	}

	if maxEndIndex < 2 {
		return nil
	}

	return &docs.Request{
		DeleteContentRange: &docs.DeleteContentRangeRequest{
			Range: &docs.Range{StartIndex: 1, EndIndex: maxEndIndex - 1},
		}}
}

func insertTableRequest(rows []TableRow) *docs.Request {
	return &docs.Request{InsertTable: &docs.InsertTableRequest{
		Location: &docs.Location{Index: 1},
		Rows:     int64(len(rows)),
		Columns:  2,
	}}
}

func createTable(srv *docs.Service, docId string, rows []TableRow) (*docs.BatchUpdateDocumentResponse, error) {
	const (
		StartIndex       = 5
		RightColumnShift = 2
		NextRowShift     = 3
	)

	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from document: %v", err)
	}

	requests := []*docs.Request{}
	clearDocReq := clearDocRequest(srv, doc)
	if clearDocReq != nil {
		requests = append(requests, clearDocReq)
	}

	requests = append(requests, insertTableRequest(rows))
	var currentLeftIndex int64 = StartIndex
	for _, row := range rows {
		text_left := row.ResponseCode
		text_right := row.Description

		requests = append(requests, &docs.Request{InsertText: &docs.InsertTextRequest{
			Text:     text_left,
			Location: &docs.Location{Index: currentLeftIndex},
		}})
		currentLeftIndex += int64(utf8.RuneCountInString(text_left)) + RightColumnShift

		requests = append(requests, &docs.Request{InsertText: &docs.InsertTextRequest{
			Text:     text_right,
			Location: &docs.Location{Index: currentLeftIndex},
		}})
		currentLeftIndex += int64(utf8.RuneCountInString(text_right)) + NextRowShift
	}

	request := docs.BatchUpdateDocumentRequest{Requests: requests}
	response, err := srv.Documents.BatchUpdate(docId, &request).Do()
	if err != nil {
		return nil, fmt.Errorf("BatchUpdate failed: %v", err)
	}

	return response, nil
}
