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
	Use:   "empty",
	Short: "Show a list of empty buckets",
	Run: func(_ *cobra.Command, _ []string) {
		s3Buddy.CheckEmpytBuckets()
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.AddCommand(emptyCmd)
}
