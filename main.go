package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
