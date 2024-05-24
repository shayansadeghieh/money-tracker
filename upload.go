package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func upload() {
	projectID := os.Getenv("PROJECT_ID")
	credentials := os.Getenv("CREDENTIALS_PATH")

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
}
