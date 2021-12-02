package cmds

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/ssmparams/pkg/ssm"
	"github.com/sirupsen/logrus"
)

// runCmgGet runs the "get" command
func runGetCmd() error {
	// Let see what we have...
	log.Debug(spew.Sdump(flags))

	// Create a new SSM client
	params, err := ssm.New(
		ssm.SetRegion(flags.awsRegion),
		ssm.SetProfile(flags.awsProfile),
	)
	if err != nil {
		return err
	}

	// Get the SSM values as listed in a []string
	outputs, err := params.GetParams(flags.param)
	if err != nil {
		return err
	}

	// Check for invalid outputs
	if len(outputs.InvalidParameters) > 0 {
		log.WithFields(logrus.Fields{
			"params": outputs.InvalidParameters,
		}).Error("parameter(s) not found")
	}

	// Dump the prams and values
	spew.Dump(outputs.Parameters)

	return nil
}
