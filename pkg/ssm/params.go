// Package ssmparams queries Amazon Web Services (AWS) Systems Manager (formerly known as SSM) Parameter Store services.
package ssm

import (
	"context"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

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

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(config.region),
		awsconfig.WithSharedConfigProfile(config.profile),
	)
	if err != nil {
		return nil, err
	}

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

// GetParam fetches the specified ssm param and returns
func (config *Config) GetParam(paramName string) (*types.Parameter, error) {
	// Fetch the value from AWS ssm.
	parameterOutput, err := config.client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: &paramName,
	})
	return parameterOutput.Parameter, err
}

func (config *Config) GetParams(paramNames []string) (map[string]interface{}, []string, error) {
	params, err := config.client.GetParameters(context.TODO(), &ssm.GetParametersInput{
		Names: paramNames,
	})
	if err != nil {
		return nil, []string{}, err
	}

	output := make(map[string]interface{}, len(params.Parameters))

	for _, v := range params.Parameters {
		output[*v.Name] = *v.Value
	}
	return output, params.InvalidParameters, nil
}

func (config *Config) PutParam(params *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	return config.client.PutParameter(context.TODO(), params)
}
