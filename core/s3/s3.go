// Package s3 is a package for s3 operations
package s3

import (
	"context"
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/olekukonko/tablewriter"
)

// isBucketEmpty checks if a bucket is empty by listing the objects in the bucket and checking if the list is empty.
// If the bucket name is empty, it will panic with the error "Bucket name is empty".
// If there is an error listing the objects, it will panic with the error "Error listing objects for bucket <bucketName>: <error>".
// Returns true if the bucket is empty, false otherwise.
func isBucketEmpty(bucketName string) bool {

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

// CheckEmpytBuckets lists all the buckets in your AWS account and checks if they are empty.
// It renders a table with the bucket name and a symbol indicating if the bucket is empty or not.
func CheckEmpytBuckets() {

	InitS3Client()

	buckets := GetAllS3SBuckets()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Bucket Name", "Empty"})

	for _, b := range buckets.Buckets {
		bucketName := aws.ToString(b.Name)

		empty := "❌"
		if isBucketEmpty(bucketName) {
			empty = "✅"
		}

		if err := table.Append([]string{bucketName, empty}); err != nil {
			panic(fmt.Sprintf("Error appending to table: %v", err))
		}
	}
	if err := table.Render(); err != nil {
		panic(fmt.Sprintf("Error rendering table: %v", err))
	}
}
