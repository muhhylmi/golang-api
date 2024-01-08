package googlesheet

import (
	"log"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"
)

func (sheetService *GSheetsServices) GetSheetData(sheetId string, sheetRange string) (*sheets.ValueRange, error) {
	values, err := sheetService.client.Spreadsheets.Values.Get(sheetId, sheetRange).Do()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (sheetService *GSheetsServices) GetSheetDataWithFilter(sheetId string, sheetRange string) (*sheets.ValueRange, []int, error) {
	filteredRows := []int{}
	filter := map[string]sheets.FilterCriteria{
		"01": {
			Condition: &sheets.BooleanCondition{
				Type: "TEXT_CONTAINS",
				Values: []*sheets.ConditionValue{
					{UserEnteredValue: "Rudi"},
				},
			},
		},
	}

	if filter != nil {
		_, err := sheetService.client.Spreadsheets.BatchUpdate(sheetId, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					SetBasicFilter: &sheets.SetBasicFilterRequest{
						Filter: &sheets.BasicFilter{
							Criteria: filter,
							Range: &sheets.GridRange{
								SheetId: 788432653,
							},
						},
					},
				},
			},
		}).Do()
		if err != nil {
			log.Fatal(err)
		}
	}

	result, err := sheetService.client.Spreadsheets.Values.Get(sheetId, sheetRange).Do()
	if err != nil {
		return nil, nil, err
	}
	fields := "sheets(data(rowMetadata(hiddenByFilter)))"
	values, err := sheetService.client.Spreadsheets.Get(sheetId).Ranges("Sheet2!A2:E").Fields(googleapi.Field(fields)).Do()
	if err != nil {
		return nil, nil, err
	}
	if filter != nil {
		rowMetadata := values.Sheets[0].Data[0].RowMetadata
		for i, r := range rowMetadata {
			// Check if "hiddenByFilter" is in the RowMetadata, it might not be present if the row is not hidden
			if !r.HiddenByFilter {
				filteredRows = append(filteredRows, i)
			}
		}

		_, err := sheetService.client.Spreadsheets.BatchUpdate(sheetId, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					ClearBasicFilter: &sheets.ClearBasicFilterRequest{
						SheetId: 788432653,
					},
				},
			},
		}).Do()
		if err != nil {
			log.Fatal(err)
		}
	}

	// return values, nil
	return result, filteredRows, nil
}
