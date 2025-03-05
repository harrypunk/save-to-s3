package client

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	err := godotenv.Load(".env.dev")
	if err != nil {
		t.Errorf("load dotenv error: %v", err)
	}

	fileUrl := os.Getenv("SRC-FILE-URL")
	endpoint := os.Getenv("DEST-S3-ENDPOINT")
	buck := os.Getenv("DEST-S3-BUCKET")
	key := os.Getenv("DEST-S3-key")
	saveClient := Client{
		Endpoint: endpoint,
	}

	ctx := context.Background()
	err = saveClient.Save(ctx, buck, key, fileUrl)
	if err != nil {
		t.Errorf("client error: %v", err)
	}
}
