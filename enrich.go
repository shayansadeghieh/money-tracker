package main

import (
	"fmt"
	"strings"
	"time"
)

func enrich(recordsJSON []AmexBill) ([]AmexBill, error) {
	for idx, record := range recordsJSON {
		recordsJSON[idx].Category = determineCategory(record.Item)
		day, month, year, err := extractTimeInformation(record.CalendarDate)
		if err != nil {
			return []AmexBill{}, err
		}

		recordsJSON[idx].Day = day
		recordsJSON[idx].Month = month
		recordsJSON[idx].Year = year

	}
	return recordsJSON, nil
}

func extractTimeInformation(dateStr string) (int, int, int, error) {
	t, err := time.Parse("01/02/2006", dateStr) //MM/DD/YYYY
	if err != nil {
		return 0, 0, 0, fmt.Errorf("unable to parse time value from string: %v", err)
	}

	// Extract the year from the time object
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	return day, month, year, nil
}

func determineCategory(item string) string {
	lowerItem := strings.ToLower(item)
	switch {
	case strings.Contains(lowerItem, "navs") || strings.Contains(lowerItem, "am2pm"):
		return "Convenience"
	case strings.Contains(lowerItem, "uber trip"):
		return "Taxi"
	case strings.Contains(lowerItem, "doordash") || strings.Contains(lowerItem, "uber eats") || strings.Contains(lowerItem, "uber one") || strings.Contains(lowerItem, "mandy's"):
		return "Takeout"
	case strings.Contains(lowerItem, "membership fee installment"):
		return "Amex Membership"
	case strings.Contains(lowerItem, "apple"):
		return "Apple"
	case strings.Contains(lowerItem, "google"):
		return "Google"
	case strings.Contains(lowerItem, "beanfield"):
		return "Internet"
	case strings.Contains(lowerItem, "disney") || strings.Contains(lowerItem, "netflix") || strings.Contains(lowerItem, "crave"):
		return "Streaming"
	case strings.Contains(lowerItem, "farmboy") || strings.Contains(lowerItem, "wal-mart") || strings.Contains(lowerItem, "metro") || strings.Contains(lowerItem, "sobeys"):
		return "Groceries"
	case strings.Contains(lowerItem, "aircanada") || strings.Contains(lowerItem, "air canada") || strings.Contains(lowerItem, "delta") || strings.Contains(lowerItem, "westjet") || strings.Contains(lowerItem, "flair"):
		return "Flights"
	case strings.Contains(lowerItem, "lcbo") || strings.Contains(lowerItem, "wine rack"):
		return "Alcohol"
	case strings.Contains(lowerItem, "gov*tor-tax"):
		return "Property Tax"
	case strings.Contains(lowerItem, "stubhub") || strings.Contains(lowerItem, "scotiabank arena"):
		return "Entertainment"
	case strings.Contains(lowerItem, "presto"):
		return "Public Transit"
	case strings.Contains(lowerItem, "maple grove"):
		return "Dentist"
	case strings.Contains(lowerItem, "amzn"):
		return "Amazon"
	case strings.Contains(lowerItem, "argonaut"):
		return "Fitness"
	case strings.Contains(lowerItem, "starbucks") || strings.Contains(lowerItem, "tim hortons") || strings.Contains(lowerItem, "coffee"):
		return "Coffee"
	default:
		return "Unknown"
	}
}
