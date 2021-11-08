package cmds

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/ssm-things/pkg/ssm"
	"github.com/sirupsen/logrus"
)

func runGetCmd() error {
	log.Debug(spew.Sdump(flags))

	params, err := ssm.New(
		ssm.SetRegion(flags.awsRegion),
		ssm.SetProfile(flags.awsProfile),
	)
	if err != nil {
		return err
	}

	ps, ip, err := params.GetParams(flags.param)
	if err != nil {
		return err
	}
	if len(ip) > 0 {
		log.WithFields(logrus.Fields{
			"params": ip,
		}).Error("parameter(s) not found")
	}
	if len(ps) > 0 {
		spew.Dump(ps)
	}

	return nil
}
