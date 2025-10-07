// Package utils is a package for utility functions
package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Ternary is a generic function that returns one of two values
// based on a boolean condition.
//
// If condition is true, it returns ifTrue, otherwise it returns ifFalse.
//
// The function is useful for writing concise and readable code
// that needs to return different values based on a condition.
func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// InitAWSClient loads the default AWS SDK config from the environment.
// It panics if there is an error loading the config.
// The function is useful for initializing the AWS SDK with the default config.
// It returns the loaded config.
func InitAWSClient() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("Unable to load SDK config, %v", err))
	}

	return cfg
}
