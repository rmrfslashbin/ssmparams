package cmds

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/ssm-things/pkg/ssm"
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
		log.Error(spew.Sdump(ip))
	}
	//log.Info(spew.Sdump(p.Parameters))

	spew.Dump(ps)

	return nil
}
