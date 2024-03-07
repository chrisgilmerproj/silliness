package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cliName     = "ssmpicker"
	ver         = "v0.0.1"
	flagVerbose = "verbose"
)

func initViper(cmd *cobra.Command) (*viper.Viper, error) {
	v := viper.New()
	errBind := v.BindPFlags(cmd.Flags())
	if errBind != nil {
		return v, fmt.Errorf("error binding flag set to viper: %w", errBind)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.SetEnvPrefix(strings.ToUpper(cliName))
	v.AutomaticEnv()
	return v, nil
}

func main() {
	rootCommand := &cobra.Command{
		Use:                   fmt.Sprintf("%s [flags]", cliName),
		DisableFlagsInUseLine: true,
		Short:                 fmt.Sprintf("%s is a CLI tool for picking ssm targets to log into", cliName),
	}

	ec2SearchCommand := &cobra.Command{
		Use:                   `ec2 [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "search ec2 instances",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  ec2Search,
	}
	initEc2SearchFlags(ec2SearchCommand.Flags())

	versionCommand := &cobra.Command{
		Use:                   `version`,
		DisableFlagsInUseLine: true,
		Short:                 "display version and quit",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  printVersion,
	}

	rootCommand.AddCommand(
		ec2SearchCommand,
		versionCommand,
	)

	if err := rootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", cliName, err.Error())
		_, _ = fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", cliName)
		os.Exit(1)
	}
}
