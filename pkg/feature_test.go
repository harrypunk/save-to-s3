package feature_test

import (
	"context"
	"github.com/harrypunk/http-to-cloud/pkg"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestS3Save(t *testing.T) {
	err := godotenv.Load("../.env.s3")
	if err != nil {
		t.Errorf("load dotenv error: %v", err)
	}

	fileUrl := os.Getenv("SRC_FILE_URL")
	endpoint := os.Getenv("DEST_S3_ENDPOINT")
	buck := os.Getenv("DEST_S3_BUCKET")
	key := os.Getenv("DEST_S3_key")
	saveClient := feature.S3Client{
		Endpoint: endpoint,
	}

	ctx := context.Background()
	err = saveClient.Save(ctx, buck, key, fileUrl)
	if err != nil {
		t.Errorf("client error: %v", err)
	}
}

func TestOssClient_Save(t *testing.T) {
	err := godotenv.Load("../.env.oss")
	if err != nil {
		t.Errorf("load dotenv error: %v", err)
	}
	fileUrl := os.Getenv("SRC_FILE_URL")
	buck := os.Getenv("DEST_OSS_BUCKET")
	key := os.Getenv("DEST_OSS_KEY")
	region := os.Getenv("OSS_REGION")

	cl := feature.OssClient{Region: region}

	err = cl.Save(context.Background(), buck, key, fileUrl)
	if err != nil {
		t.Errorf("failed to save oss: %v", err)
	}
}

func TestCosSave1(t *testing.T) {
	err := godotenv.Load("../.env.cos.put")
	if err != nil {
		t.Errorf("load dotenv error: %v", err)
	}
	fileUrl := os.Getenv("SRC_FILE_URL")
	endpoint := os.Getenv("DEST_COS_ENDPOINT")
	key := os.Getenv("DEST_COS_KEY")
	tcId := os.Getenv("TC_ID")
	tcKey := os.Getenv("TC_KEY")

	cl := feature.CosPutClient{Endpoint: endpoint, TCId: tcId, TCKey: tcKey}
	err = cl.Save(context.Background(), "", key, fileUrl)
	if err != nil {
		t.Errorf("failed to save oss: %v", err)
	}
}
