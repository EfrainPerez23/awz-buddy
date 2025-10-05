// Package s3 is a package for s3 operations
package s3

import (
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/olekukonko/tablewriter"

	UTILS "awz-buddy/core"
)

// CheckPublicBuckets lists all the buckets in your AWS account and checks the status of "Block all public access" and "Ignore all public ACLs".
// It renders a table with the bucket name and four columns indicating the status of the two policies.
// If a bucket name is provided, only that bucket is checked, otherwise all the buckets are checked.
// If an error occurs, the program will panic with the error message.
func CheckPublicBuckets(bucketName *string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Bucket Name", "BlockPublicAcls", "IgnorePublicAcls", "BlockPublicPolicy", "RestrictPublicBuckets"})

	if bucketName != nil {
		generateContentPolicyTable(*bucketName, table)
	} else {
		buckets := GetAllS3SBuckets()
		for _, bucket := range buckets.Buckets {
			generateContentPolicyTable(*bucket.Name, table)
		}
	}

	if err := table.Render(); err != nil {
		panic(fmt.Sprintf("Error rendering table: %v", err))
	}
}

// generateContentPolicyTable generates a table with the bucket name and four columns indicating the status of the two policies.
// The columns are:
// - BlockPublicAcls: indicates if the bucket has the "Block public ACLs" policy enabled.
// - IgnorePublicAcls: indicates if the bucket has the "Ignore public ACLs" policy enabled.
// - BlockPublicPolicy: indicates if the bucket has the "Block public policy" policy enabled.
// - RestrictPublicBuckets: indicates if the bucket has the "Restrict public buckets" policy enabled.
// If an error occurs while generating the table, the program will panic with the error message.
func generateContentPolicyTable(bucketName string, table *tablewriter.Table) {
	blockAccess := CheckPublicPoliciesForBuckets(bucketName).PublicAccessBlockConfiguration

	appendErr := table.Append([]string{
		bucketName,
		fmt.Sprintf("%v", UTILS.Ternary(aws.ToBool(blockAccess.BlockPublicAcls), "✅", "❌")),
		fmt.Sprintf("%v", UTILS.Ternary(aws.ToBool(blockAccess.IgnorePublicAcls), "✅", "❌")),
		fmt.Sprintf("%v", UTILS.Ternary(aws.ToBool(blockAccess.BlockPublicPolicy), "✅", "❌")),
		fmt.Sprintf("%v", UTILS.Ternary(aws.ToBool(blockAccess.RestrictPublicBuckets), "✅", "❌")),
	})

	if appendErr != nil {
		panic(fmt.Sprintf("Error appending to table: %v", appendErr))
	}
}
