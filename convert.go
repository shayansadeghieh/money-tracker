package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AmexBill struct {
	CalendarDate string
	ID           string
	Amount       float32
	Item         string
	Category     string
	Day          int
	Month        int
	Year         int
}

func removeWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func convertCSVToJSON(records [][]string) ([]AmexBill, error) {
	var amexBillSection AmexBill
	var amexBill []AmexBill

	for _, val := range records {
		parsedAmount, err := strconv.ParseFloat(removeWhitespace(val[2]), 32)
		if err != nil {
			return []AmexBill{}, fmt.Errorf("unable to parse amount from records: %v", err)
		}
		// Only consider things purchased in toronto and that actually cost money
		if strings.Contains(strings.ToLower(val[3]), "toronto") && float32(parsedAmount) > 0 {
			amexBillSection.CalendarDate = removeWhitespace(val[0])
			amexBillSection.ID = removeWhitespace(val[1])
			amexBillSection.Amount = float32(parsedAmount)

			amexBillSection.Item = removeWhitespace(val[3])
			amexBill = append(amexBill, amexBillSection)

		}

	}
	return amexBill, nil
}
