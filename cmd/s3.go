// Package cmd is a package for the root command
package cmd

import (
	s3Buddy "awz-buddy/core/s3"

	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "s3 commands",
}

var emptyCmd = &cobra.Command{
	Use:   "empty [bucket_name]",
	Short: "List all S3 buckets that are empty",
	PreRun: func(cmd *cobra.Command, _ []string) {
		s3Buddy.InitS3Client()
	},
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 1 {
			bucketName := args[0]
			s3Buddy.CheckEmpytBuckets(&bucketName)
			return
		}
		s3Buddy.CheckEmpytBuckets(nil)
	},
}

var publicCmd = &cobra.Command{
	Use: "public [bucket_name]",

	Short: "Check whether S3 buckets have 'Block all public access' enabled (✅) or disabled (❌)",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, _ []string) {
		s3Buddy.InitS3Client()
	},
	Run: func(_ *cobra.Command, args []string) {

		if len(args) == 1 {
			bucketName := args[0]
			s3Buddy.CheckPublicBuckets(&bucketName)
			return
		}
		s3Buddy.CheckPublicBuckets(nil)
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.AddCommand(emptyCmd)
	s3Cmd.AddCommand(publicCmd)
}
