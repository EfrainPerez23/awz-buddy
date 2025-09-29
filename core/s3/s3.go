package main

import (
	"context"
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/olekukonko/tablewriter"
)

var s3Client *s3.Client

func isBucketEmpty(bucketName string) bool {

	if bucketName == "" {
		panic("Bucket name is empty")
	}

	output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		panic(fmt.Sprintf("Error listing objects for bucket %s: %v", bucketName, err))
	}
	return len(output.Contents) == 0
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	s3Client = s3.NewFromConfig(cfg)

	buckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		panic(fmt.Sprintf("Error listing buckets: %v", err))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Bucket Name", "Empty"})
	for _, b := range buckets.Buckets {
		bucketName := aws.ToString(b.Name)

		empty := "❌"
		if isBucketEmpty(bucketName) {
			empty = "✅"
		}

		table.Append([]string{bucketName, empty})
	}
	table.Render()
}
