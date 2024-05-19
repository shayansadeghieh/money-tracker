package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		fmt.Println("Error reading records")
	}

	csvMap := make(map[string]int)
	csvMap["item"] = 3

	categorizedRecords := addCategory(records, csvMap)
	fmt.Println(categorizedRecords)

}

func addCategory(records [][]string, csvMap map[string]int) [][]string {
	for idx := range records {
		records[idx] = append(records[idx], determineCategory(records[idx][csvMap["item"]]))
	}
	return records
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
	case strings.Contains(lowerItem, "disney") || strings.Contains(lowerItem, "netflix"):
		return "Streaming"
	case strings.Contains(lowerItem, "farmboy") || strings.Contains(lowerItem, "metro") || strings.Contains(lowerItem, "sobeys"):
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
	case strings.Contains(lowerItem, "starbucks"):
		return "Coffee"
	default:
		return "Unknown"
	}
}
