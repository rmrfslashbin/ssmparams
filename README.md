# ssmparams
[ssmparams](https://github.com/rmrfslashbin/ssmparams) provides a simple method to asynchronosly fetch parameters from the [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html).

[![Go](https://github.com/rmrfslashbin/ssmparams/actions/workflows/go.yml/badge.svg)](https://github.com/rmrfslashbin/ssmparams/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rmrfslashbin/ssmparams)](https://goreportcard.com/report/github.com/rmrfslashbin/ssmparams)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rmrfslashbin/ssmparams)
[![Go Reference](https://pkg.go.dev/badge/github.com/rmrfslashbin/ssmparams.svg)](https://pkg.go.dev/github.com/rmrfslashbin/ssmparams)

## Configuration
This module expects a configured AWS credentials file with a default profile. See https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html for more details.

## Example
```
package main

import (
    "fmt"
    "github.com/rmrfslashbin/ssmparams"
)

func main() {
    // Create new SSMParams struct.
	ssmps := ssmparams.SSMParams{}
	if err := ssmps.New(); err != nil {
		// Bail out on error.
        panic(err)
	}

    // Fetch a named paramerter from the Parameter Store.
    // Returns a channel while fetching data.
	ch := ssmps.GetParam("/some/nifty/param")

    // Block, waiting for channel to return data.
	clientID := <-ch
	if clientID.Err != nil {
		panic(clientID.Err)
	}

    // Print data and values fetehed from the Parameter store.
	fmt.Printf("Name:               %v\n", *clientID.ParameterOutput.Parameter.Name)
	fmt.Printf("Value:              %v\n", *clientID.ParameterOutput.Parameter.Value)
	fmt.Printf("Type:               %v\n", clientID.ParameterOutput.Parameter.Type)
	fmt.Printf("Last Modified Date: %v\n", clientID.ParameterOutput.Parameter.LastModifiedDate)
}
```
