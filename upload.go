package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func upload(records []AmexBill) error {
	projectID := os.Getenv("PROJECT_ID")
	datasetID := "amex_bills"
	tableID := "shayan"

	credentials := os.Getenv("CREDENTIALS_PATH")

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(credentials))
	if err != nil {
		return fmt.Errorf("failed to create bigquery client %v", err)
	}
	defer client.Close()

	// Get the table reference.
	table := client.Dataset(datasetID).Table(tableID)

	// Upload the data.
	inserter := table.Inserter()
	if err := inserter.Put(ctx, records); err != nil {
		return fmt.Errorf("failed to insert rows %v", err)
	}

	return nil
}
