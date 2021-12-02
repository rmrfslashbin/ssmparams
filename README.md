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
	// awscli profile name
	awsProfile := "default"
	awsRegion := "us-east-1"
	
    params, err := ssm.New(
		ssm.SetProfile(awsProfile),
		ssm.SetRegion(awsRegion),
	)
	if err != nil {
		panic(err)
	}

	mapOfReturnParams, InvalidParameters, err := params.GetParams(flags.param)
	if err != nil {
		return err
	}
	if len(InvalidParameters) > 0 {
		log.WithFields(logrus.Fields{
			"params": InvalidParameters,
		}).Error("parameter(s) not found")
	}
	if len(mapOfReturnParams) > 0 {
		spew.Dump(mapOfReturnParams)
	}
}
```
