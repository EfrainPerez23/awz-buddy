// Package s3 is a package for s3 operations

package s3

import (
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/olekukonko/tablewriter"

	UTILS "awz-buddy/core"
)

// CheckEmpytBuckets lists all the buckets in your AWS account and checks if they are empty.
// It renders a table with the bucket name and a symbol indicating if the bucket is empty or not.
func CheckEmpytBuckets(bucketName *string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Bucket Name", "Empty"})

	if bucketName != nil {
		generateContentEmptyTable(*bucketName, table)
	} else {
		buckets := GetAllS3SBuckets()

		for _, b := range buckets.Buckets {
			bucketName := aws.ToString(b.Name)

			generateContentEmptyTable(bucketName, table)
		}
	}

	if err := table.Render(); err != nil {
		panic(fmt.Sprintf("Error rendering table: %v", err))
	}
}

// generateContentEmptyTable generates a table row for the given bucket name, with a column indicating if the bucket is empty or not.
// The first column is the bucket name, and the second column is either "✅" (empty) or "❌" (not empty).
// If an error occurs while appending to the table, the program will panic with the error message.
func generateContentEmptyTable(bucketName string, table *tablewriter.Table) {
	appendErr := table.Append([]string{bucketName, UTILS.Ternary(IsBucketEmpty(bucketName), "✅", "❌")})
	if appendErr != nil {
		panic(fmt.Sprintf("Error appending to table: %v", appendErr))
	}
}
