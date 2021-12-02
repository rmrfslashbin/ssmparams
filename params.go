// Package ssmparams queries Amazon Web Services (AWS) Systems Manager (formerly known as SSM) Parameter Store services.
package ssmparams

import (
	"context"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// ParamOutput is the output of the GetParam function.
type ParamOutput struct {
	// Parameters is a map of parameter names to values.
	Parameters map[string]interface{}

	// InvalidParameters is a list of parameter names that were not found.
	InvalidParameters []string
}

// Used to manage varidic options
type Option func(config *Config)

// Config is used to configure the ssmparams client.
type Config struct {
	// region is the AWS region to use.
	region string

	// profile is the AWS profile to use.
	profile string

	// client is the AWS ssm client to use.
	client *ssm.Client
}

// New is a factory function for creating a new Config
func New(opts ...func(*Config)) (*Config, error) {
	config := &Config{}

	// Set default values
	config.profile = "default"
	config.region = "us-east-1"

	// apply the list of options to Config
	for _, opt := range opts {
		opt(config)
	}

	// Create the AWS config
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(config.region),
		awsconfig.WithSharedConfigProfile(config.profile),
	)
	if err != nil {
		return nil, err
	}

	// Create the AWS ssm client
	config.client = ssm.NewFromConfig(cfg)

	return config, nil
}

// SetRegion sets the AWS region to use.
func SetRegion(region string) Option {
	return func(config *Config) {
		config.region = region
	}
}

// SetProfile sets the AWS profile to use.
func SetProfile(profile string) Option {
	return func(config *Config) {
		config.profile = profile
	}
}

// GetParam gets a parameter from AWS Parameter Store.
func (config *Config) GetParams(paramNames []string) (*ParamOutput, error) {
	// Fetch the requested parameters
	params, err := config.client.GetParameters(context.TODO(), &ssm.GetParametersInput{
		Names: paramNames,
	})

	// Bail out on error
	if err != nil {
		return nil, err
	}

	// Setup a map to store the valid parameters
	output := make(map[string]interface{}, len(params.Parameters))

	// Loop through the parameters and store them in the map
	for _, v := range params.Parameters {
		output[*v.Name] = *v.Value
	}

	// Done!
	return &ParamOutput{Parameters: output, InvalidParameters: params.InvalidParameters}, nil
}

// PutParam puts a parameter into AWS Parameter Store.
func (config *Config) PutParam(params *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	return config.client.PutParameter(context.TODO(), params)
}
