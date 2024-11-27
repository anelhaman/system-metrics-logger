package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// writeToGoogleSheet writes metrics to a sheet with the hostname as the sheet name.
func writeToGoogleSheet(sheetID string, values [][]interface{}) error {
	// Load credentials
	credFile, err := os.ReadFile("credentials.json")
	if err != nil {
		return fmt.Errorf("unable to read credentials file: %w", err)
	}

	// Create a new Sheets service
	ctx := context.Background()
	config, err := google.JWTConfigFromJSON(credFile, sheets.SpreadsheetsScope)
	if err != nil {
		return fmt.Errorf("unable to parse credentials file: %w", err)
	}
	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to create Sheets service: %w", err)
	}

	// Get the hostname to use as the sheet name
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("unable to get hostname: %w", err)
	}

	// Get the spreadsheet to check sheet existence
	spreadsheet, err := srv.Spreadsheets.Get(sheetID).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve spreadsheet: %w", err)
	}

	var sheetIDForHostname int64
	sheetExists := false
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == hostname {
			sheetExists = true
			sheetIDForHostname = sheet.Properties.SheetId
			break
		}
	}

	// Create the sheet if it doesn't exist
	if !sheetExists {
		_, err := srv.Spreadsheets.BatchUpdate(sheetID, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title: hostname,
						},
					},
				},
			},
		}).Do()
		if err != nil {
			return fmt.Errorf("unable to create sheet: %w", err)
		}
	}

	// Write headers if sheet is empty or first row is empty
	headerRange := fmt.Sprintf("%s!A1:D1", hostname)
	headers := []interface{}{"Timestamp", "CPU Usage (%)", "Memory Usage (%)", "Disk Usage (%)"}

	_, err = srv.Spreadsheets.Values.Update(sheetID, headerRange, &sheets.ValueRange{
		Values: [][]interface{}{headers},
	}).ValueInputOption("RAW").Do()

	if err != nil {
		return fmt.Errorf("unable to write headers: %w", err)
	}

	// Write values (append data)
	dataRange := fmt.Sprintf("%s!A2:D", hostname)
	_, err = srv.Spreadsheets.Values.Append(sheetID, dataRange, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("RAW").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return fmt.Errorf("unable to append data: %w", err)
	}

	// Apply formatting (center alignment)
	_, err = srv.Spreadsheets.BatchUpdate(sheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				RepeatCell: &sheets.RepeatCellRequest{
					Range: &sheets.GridRange{
						SheetId:          sheetIDForHostname,
						StartRowIndex:    0,
						EndRowIndex:      int64(len(values) + 1), // Include header row
						StartColumnIndex: 0,
						EndColumnIndex:   4, // A-D columns
					},
					Cell: &sheets.CellData{
						UserEnteredFormat: &sheets.CellFormat{
							HorizontalAlignment: "CENTER",
						},
					},
					Fields: "userEnteredFormat(horizontalAlignment)",
				},
			},
		},
	}).Do()
	if err != nil {
		return fmt.Errorf("unable to apply formatting: %w", err)
	}

	return nil
}
