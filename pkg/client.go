package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	Endpoint string
}

func (cl *Client) UploadFileToS3(ctx context.Context, bucketName, objectKey, fileURL string) error {
	var cfg aws.Config
	var err error

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cl.Endpoint)
	})

	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("http get error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   resp.Body,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}

	fmt.Printf("Successfully uploaded %s to s3://%s/%s\n", fileURL, bucketName, objectKey)
	return nil
}
