package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

type SpreadsheetPushRequest struct {
	SpreadsheetId string        `json:"spreadsheet_id"`
	Range         string        `json:"range"`
	Values        []interface{} `json:"values"`
}

type SpreadsheetService struct {
	service *sheets.Service
}

const (
	client_secret_path = "D:\\go-web-socket\\src\\secret.json"
)

func NewSpreadsheetService() (*SpreadsheetService, error) {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile(client_secret_path), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		log.Fatal("GoogleClientErr : ", err)
	}
	return &SpreadsheetService{service: srv}, nil
}

func (s *SpreadsheetService) WriteToSpreadsheet(object *SpreadsheetPushRequest) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, object.Values)
	_, err := s.service.Spreadsheets.Values.Update(object.SpreadsheetId, object.Range, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		fmt.Println("SheetWriteErr : ", err)
	}
	return err
}
