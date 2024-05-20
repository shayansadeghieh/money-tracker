package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
		amexBillSection.CalendarDate = removeWhitespace(val[0])
		amexBillSection.ID = removeWhitespace(val[1])

		parsedAmount, err := strconv.ParseFloat(removeWhitespace(val[2]), 32)
		if err != nil {
			return []AmexBill{}, fmt.Errorf("unable to parse amount from records: %v", err)
		}
		amexBillSection.Amount = float32(parsedAmount)

		amexBillSection.Item = removeWhitespace(val[3])
		amexBill = append(amexBill, amexBillSection)
	}
	return amexBill, nil
}

func main() {
	path := "/Users/shayansadeghieh/amex-bills/ofx.csv"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records from CSV")
	}

	recordsJSON, err := convertCSVToJSON(records)
	if err != nil {
		log.Fatalf("Error converting CSV to JSON: %v", err)
	}

	enrichRecords, err := enrich(recordsJSON)
	if err != nil {
		log.Fatalf("Error enriching records: %v", err)
	}
	fmt.Println(enrichRecords)

}

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
	t, err := time.Parse("01/02/2006", dateStr)
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
