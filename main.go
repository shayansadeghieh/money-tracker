package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	path := "/Users/shayansadeghieh/amex-bills/ofx.csv"

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error while reading the file %v:", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading records from CSV %v", err)
	}

	recordsStruct, err := convertCSVToStruct(records)
	if err != nil {
		log.Fatalf("Error converting CSV to Struct: %v", err)
	}

	enrichedRecords, err := enrich(recordsStruct)
	if err != nil {
		log.Fatalf("Error enriching records: %v", err)
	}

	err = upload(enrichedRecords)
	if err != nil {
		log.Fatalf("Failed to upload data to BigQuery: %v", err)
	}

}
