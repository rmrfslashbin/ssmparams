package cmds

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Flags struct contains settings for the root command
type Flags struct {
	loglevel   string
	json       bool
	pretty     bool
	confighome string
	param      []string
	awsRegion  string
	awsProfile string
}

var (
	flags Flags
	log   *logrus.Logger

	// rootCmd is the Viper root command
	rootCmd = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set the log level
			switch flags.loglevel {
			case "error":
				log.SetLevel(logrus.ErrorLevel)
			case "warn":
				log.SetLevel(logrus.WarnLevel)
			case "info":
				log.SetLevel(logrus.InfoLevel)
			case "debug":
				log.SetLevel(logrus.DebugLevel)
			case "trace":
				log.SetLevel(logrus.TraceLevel)
			default:
				log.SetLevel(logrus.InfoLevel)
			}

			if flags.json {
				log.SetFormatter(&logrus.JSONFormatter{})
				if flags.pretty {
					log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
				}
			}
		},
	}

	cmdGet = &cobra.Command{
		Use:   "get",
		Short: "get param",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runGetCmd(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		},
	}
)

// init sets up the CLI and flags
func init() {
	// Set the log level
	log = logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// get platform specific user config directory
	flags.confighome, _ = os.UserConfigDir()

	rootCmd.PersistentFlags().StringVarP(&flags.loglevel, "loglevel", "", "info", "[error|warn|info|debug|trace]")
	rootCmd.PersistentFlags().BoolVarP(&flags.pretty, "pretty", "", false, "pretty print the json")
	rootCmd.PersistentFlags().BoolVarP(&flags.json, "json", "", false, "output as json")
	rootCmd.PersistentFlags().StringVarP(&flags.awsProfile, "profile", "", "default", "AWS profile")
	rootCmd.PersistentFlags().StringVarP(&flags.awsRegion, "region", "", "us-east-1", "AWS profile")

	cmdGet.PersistentFlags().StringSliceVarP(&flags.param, "param", "p", []string{}, "param to get (required)")
	cmdGet.MarkPersistentFlagRequired("param")

	rootCmd.AddCommand(
		cmdGet,
	)

}

// Execute the root command
func Execute() error {
	return rootCmd.Execute()
}
