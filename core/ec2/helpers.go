// Package ec2 is a package for ec2 operations
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"

	CORE "awz-buddy/core"
)

var Ec2Client *ec2.Client

// InitEc2Client initializes the EC2 client with the default AWS config.
// It panics if there is an error loading the SDK config.
// The function is useful for initializing the EC2 client with the default config.
// It returns the loaded config.
func InitEc2Client() {

	if Ec2Client != nil {
		return
	}

	cfg := CORE.InitAWSClient()

	Ec2Client = ec2.NewFromConfig(cfg)
}

// GetAllElasticIps lists all the Elastic IP addresses in your AWS account.
// It panics if there is an error listing the addresses.
// Returns a pointer to the *ec2.DescribeAddressesOutput response.
// The response contains a list of Elastic IP addresses along with their allocation IDs and public IPs.

func GetAllElasticIps() *ec2.DescribeAddressesOutput {
	eips, err := Ec2Client.DescribeAddresses(context.TODO(), &ec2.DescribeAddressesInput{})

	if err != nil {
		panic(fmt.Sprintf("Error getting all EIPs: %v\n", err))
	}

	return eips
}
