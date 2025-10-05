// Package s3 is a package for s3 operations
package s3

import (
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/olekukonko/tablewriter"
)

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
		if IsBucketEmpty(bucketName) {
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
