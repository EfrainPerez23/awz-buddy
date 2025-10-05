// Package s3 is a package for s3 operations
package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

// InitS3Client initializes the S3 client with the default AWS config.
// It panics if there is an error loading the SDK config.
func InitS3Client() {

	if S3Client != nil {
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("Unable to load SDK config, %v", err))
	}

	// Create a new S3 client from the loaded config
	S3Client = s3.NewFromConfig(cfg)
}

// IsBucketEmpty checks if a bucket is empty by listing the objects in the bucket and checking if the list is empty.
// If the bucket name is empty, it will panic with the error "Bucket name is empty".
// If there is an error listing the objects, it will panic with the error "Error listing objects for bucket <bucketName>: <error>".
// Returns true if the bucket is empty, false otherwise.
func IsBucketEmpty(bucketName string) bool {

	InitS3Client()

	if bucketName == "" {
		panic("Bucket name is empty")
	}

	output, err := S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		panic(fmt.Sprintf("Error listing objects for bucket %s: %v", bucketName, err))
	}
	return len(output.Contents) == 0
}

// GetAllS3SBuckets lists all the buckets in your AWS account and returns a pointer to the *s3.ListBucketsOutput response.
// It panics if there is an error listing the buckets.
func GetAllS3SBuckets() *s3.ListBucketsOutput {

	InitS3Client()
	allBuckets, err := S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		panic(fmt.Sprintf("Error listing buckets: %v", err))
	}

	return allBuckets
}
