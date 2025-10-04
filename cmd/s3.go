package cmd

import (
	s3Buddy "awz-buddy/core/s3"

	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Comandos para S3",
}

var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Show a list of empty buckets",
	Run: func(_ *cobra.Command, args []string) {
		s3Buddy.CheckEmpytBuckets()
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.AddCommand(emptyCmd)
}
