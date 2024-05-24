package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestEnrich(t *testing.T) {
	testCases := []struct {
		name         string
		inputRecords []AmexBill
		expRecords   []AmexBill
		expError     error
	}{
		{
			name: "Normal records",
			inputRecords: []AmexBill{
				{
					CalendarDate: "05/07/2023",
					ID:           "123",
					Amount:       45.54,
					Item:         "air canada",
					Category:     "",
					Day:          0,
					Month:        0,
					Year:         0,
				},
			},
			expRecords: []AmexBill{
				{
					CalendarDate: "05/07/2023",
					ID:           "123",
					Amount:       45.54,
					Item:         "air canada",
					Category:     "Flights",
					Day:          7,
					Month:        5,
					Year:         2023,
				},
			},
			expError: nil,
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			records, err := enrich(tc.inputRecords)
			if err == nil && tc.expError != nil || err != nil && tc.expError == nil {
				t.Errorf("Handling error: Got %v, but expected %v", tc.expError, err)
			}
			if !reflect.DeepEqual(records, tc.expRecords) {
				t.Errorf("Error enriching records: Got %v, but expected %v", records, tc.expRecords)
			}

		})
	}
}

func TestExtractTimeInformation(t *testing.T) {
	testCases := []struct {
		name     string
		dateStr  string
		expDay   int
		expMonth int
		expYear  int
		expError error
	}{
		{
			name:     "Normal date",
			dateStr:  "07/05/2020",
			expDay:   5,
			expMonth: 7,
			expYear:  2020,
			expError: nil,
		},
		{
			name:     "Incorrect date format",
			dateStr:  "07-05-2020",
			expDay:   0,
			expMonth: 0,
			expYear:  0,
			expError: errors.New("unable to parse time value from string"),
		},
		{
			name:     "Incomplete date",
			dateStr:  "07/05/202",
			expDay:   0,
			expMonth: 0,
			expYear:  0,
			expError: errors.New("unable to parse time value from string"),
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			day, month, year, err := extractTimeInformation(tc.dateStr)

			if err == nil && tc.expError != nil || err != nil && tc.expError == nil {
				t.Errorf("Handling error: Got %v, but expected %v", err, tc.expError)
			}

			if !reflect.DeepEqual(day, tc.expDay) {
				t.Errorf("Extracting day: Got %v, but expected %v", day, tc.expDay)
			}

			if !reflect.DeepEqual(month, tc.expMonth) {
				t.Errorf("Extracting month: Got %v, but expected %v", month, tc.expMonth)
			}

			if !reflect.DeepEqual(year, tc.expYear) {
				t.Errorf("Extracting year: Got %v, but expected %v", year, tc.expYear)
			}

		})
	}
}
