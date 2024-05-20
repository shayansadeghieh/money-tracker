package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertCSVToJSON(t *testing.T) {
	testCases := []struct {
		name      string
		records   [][]string
		expOutput []AmexBill
		expError  error
	}{
		{
			name: "Normal records",
			records: [][]string{{
				"01/03/2024", "1234", "278.6", "shoes",
			}},
			expOutput: []AmexBill{
				{
					CalendarDate: "01/03/2024",
					ID:           "1234",
					Amount:       278.6,
					Item:         "shoes",
				},
			},
			expError: nil,
		},
		{
			name: "Incomplete records",
			records: [][]string{{
				"01/03/2024", "278.6", "shoes",
			}},
			expOutput: []AmexBill{},
			expError:  fmt.Errorf("Incomplete records. We expect an error."),
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			recordsJSON, err := convertCSVToJSON(tc.records)
			if err == nil && tc.expError != nil || err != nil && tc.expError == nil {
				t.Errorf("Error when handling non-nil error: Got %v, but expected %v", tc.expError, err)
			}
			if !reflect.DeepEqual(recordsJSON, tc.expOutput) {
				t.Errorf("Error converting records CSV to JSON: Got %v, but expected %v", recordsJSON, tc.expOutput)
			}

		})
	}
}
