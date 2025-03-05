package client

import (
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	fileUrl := os.Getenv("SRC-FILE-URL")
	endpoint := os.Getenv("DEST-S3-ENDPOINT")
	buck := os.Getenv("DEST-S3-BUCKET")
	key := os.Getenv("DEST-S3-key")
	saveClient := Client{
		Endpoint: endpoint,
	}

	err := saveClient.Save(buck, key, fileUrl)
	if err != nil {
		t.Errorf("client error: %v", err)
	}
}
