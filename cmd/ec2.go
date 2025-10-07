// Package cmd is a package for the root command
package cmd

import (
	"github.com/spf13/cobra"

	EC2 "awz-buddy/core/ec2"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "ec2 commands",
}

var eipCmd = &cobra.Command{
	Use:   "eip",
	Short: "List all Elastic IP addresses that are used or unused",
	PreRun: func(cmd *cobra.Command, _ []string) {
		EC2.InitEc2Client()
	},
	Run: func(_ *cobra.Command, _ []string) {
		EC2.CheckElasticIps()
	},
}

func init() {
	rootCmd.AddCommand(ec2Cmd)
	ec2Cmd.AddCommand(eipCmd)
}
