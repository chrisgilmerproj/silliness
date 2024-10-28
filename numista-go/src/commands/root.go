package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	CliName = "numista"

	// Global Flags
	flagLang = "lang"

	// Languages
	langEn = "en"
	langEs = "es"
	langFr = "fr"
)

var (
	langChoices = []string{langEn, langEs, langFr}
)

func initViper(cmd *cobra.Command) (*viper.Viper, error) {
	v := viper.New()
	errBind := v.BindPFlags(cmd.Flags())
	if errBind != nil {
		return v, fmt.Errorf("error binding flag set to viper: %w", errBind)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.SetEnvPrefix(strings.ToUpper(CliName))
	v.AutomaticEnv()
	return v, nil
}

// Initialize the flags for the root command
func initRootFlags(flag *pflag.FlagSet) {
	flag.String(flagLang, langEn, "Set the language")
}

func CreateCommands(version string) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:                   fmt.Sprintf("%s [flags]", CliName),
		DisableFlagsInUseLine: true,
		Short:                 fmt.Sprintf("%s is a CLI tool for working with the API", CliName),
	}
	initRootFlags(rootCommand.PersistentFlags())

	mintsCommand := &cobra.Command{
		Use:                   `mints [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list mints",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  mints,
	}
	initMintsFlags(mintsCommand.Flags())

	searchTypesCommand := &cobra.Command{
		Use:                   `search-types [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "search types",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  searchTypes,
	}
	initSearchTypesFlags(searchTypesCommand.Flags())

	versionCommand := &cobra.Command{
		Use:                   `version`,
		DisableFlagsInUseLine: true,
		Short:                 "display the version and quit",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s version %s\n", CliName, version)
			return nil
		},
	}

	rootCommand.AddCommand(
		mintsCommand,
		searchTypesCommand,
		versionCommand,
	)
	return rootCommand
}
