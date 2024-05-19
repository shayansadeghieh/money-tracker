package main

import (
	"context"
	"os"

	"cloud.google.com/go/bigquery"
)

func upload(path string) error {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		// TODO: Handle error.
	}

}
