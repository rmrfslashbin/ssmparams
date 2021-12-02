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
	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/ssmparams"
	"github.com/sirupsen/logrus"
)

func main() {
	// awscli profile name
	awsProfile := "default"
	awsRegion := "us-east-1"

	// Set up a new ssmparams client
	params, err := ssmparams.New(
		ssmparams.SetProfile(awsProfile),
		ssmparams.SetRegion(awsRegion),
	)
	if err != nil {
		panic(err)
	}

	// Set up a string slice of parameter names to retrieve
	request := []string{
		"/first/param/to/request/secret",
		"/first/param/to/request/secret",
		"/second/param/to/request",
	}

	// Retrieve the parameters
	outputs, err := params.GetParams(request)
	if err != nil {
		panic(err)
	}

	// Check if invalid parameters were returned
	if len(outputs.InvalidParameters) > 0 {
		logrus.WithFields(logrus.Fields{
			"params": outputs.InvalidParameters,
		}).Error("parameter(s) not found")
	}

	// Dump the output
	spew.Dump(outputs.Parameters)
}
```
