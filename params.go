// Amazon Web Services (AWS) Systems Manager (formerly known as SSM) Parameter Store services.
package ssmparams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// ParamOutput struct holds output values and errors.
type ParamOutput struct {
	ParameterOutput *ssm.GetParameterOutput
	Err             error
}

// SSMParams struct hold the SSM client
// and other operational data.
type SSMParams struct {
	client *ssm.Client
}

// New sets up AWS auth and binds to an SSM client.
func New() (*SSMParams, error) {
	ssmparams := &SSMParams{}

	// Load config from ENV or .aws credentials file.
	// TODO: Add funtionality for different auth methods.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return ssmparams, err
	}

	// Create a new ssm client.
	ssmparams.client = ssm.NewFromConfig(cfg)
	return ssmparams, nil
}

// GetParam fetches the specified ssm param and returns
// a channel to fetch the value asynchronously.
func (s *SSMParams) GetParam(paramName string) <-chan ParamOutput {
	// Create a channel to return the value.
	c := make(chan ParamOutput)

	// Goroutine to fetch the value.
	go func() {
		defer close(c)

		// Set up the request.
		input := &ssm.GetParameterInput{
			Name: &paramName,
		}

		// Fetch the value from AWS ssm.
		parameterOutput, err := s.client.GetParameter(context.TODO(), input)
		// Send the value to the channel.
		c <- ParamOutput{parameterOutput, err}
	}() // execute the goroutine

	// Return the channel.
	return c
}
