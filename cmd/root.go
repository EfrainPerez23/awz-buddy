// Package cmd is a package for the root command
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "awz-buddy",
	Short: "AWZ Buddy is a CLI tool to help you manage and audit your AWS resources",
	Long:  `AWZ Buddy allows you to detect unused or misconfigured resources, prevent cost leaks, and keep your AWS environment clean and secure`,
}

// Execute runs the root command and prints any errors to the console.
// If an error occurs, the program will exit with a status code of 1.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
