// Package ec2 is a package for ec2 operations
package ec2

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/olekukonko/tablewriter"
)

// CheckElasticIps lists all the Elastic IP addresses in your AWS account and checks if they are associated with an instance or not.
// It renders a table with the allocation ID, public IP, association instance ID, and associated ID.
// If an error occurs while rendering the table, the program will panic with the error message.
func CheckElasticIps() {
	eips := GetAllElasticIps()
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Public Ip", "Allocation Id", "Association Instance Id", "Associated Id"})

	for _, eip := range eips.Addresses {

		publicIP := aws.ToString(eip.PublicIp)
		allocationID := aws.ToString(eip.AllocationId)

		if eip.AssociationId == nil {
			appendEipTable(table, publicIP, allocationID, "UNUSED EIP 💡", "UNUSED EIP 💡")
			continue
		}

		assoc := aws.ToString(eip.AssociationId)
		instance := aws.ToString(eip.InstanceId)
		eni := aws.ToString(eip.NetworkInterfaceId)

		target := instance
		if target == "" {
			target = eni
		}

		appendEipTable(table, publicIP, allocationID, target, assoc)

	}

	if err := table.Render(); err != nil {
		panic(fmt.Sprintf("Error rendering table: %v", err))
	}

}

func appendEipTable(table *tablewriter.Table, publicIp string, allocationId string, target string, assocId string) {
	appendErr := table.Append([]string{
		publicIp,
		allocationId,
		target,
		assocId,
	})
	if appendErr != nil {
		panic(fmt.Sprintf("Error appending to table: %v", appendErr))
	}
}
