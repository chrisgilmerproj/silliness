package commands

import (
	"fmt"
	"strings"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/utils"
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

func validateRootFlags(v *viper.Viper) error {
	lang := v.GetString(flagLang)
	if !utils.Contains(langChoices, lang) {
		return fmt.Errorf("lang must be one of %v", langChoices)
	}
	return nil
}

func CreateCommands(version string) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:                   fmt.Sprintf("%s [flags]", CliName),
		DisableFlagsInUseLine: true,
		Short:                 fmt.Sprintf("%s is a CLI tool for working with the API", CliName),
	}
	initRootFlags(rootCommand.PersistentFlags())

	typesSubcommand := &cobra.Command{
		Use:                   `types [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "types commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	getTypeCommand := &cobra.Command{
		Use:                   `get [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get types",
		Long:                  "Retrieve the details about a specific type in the catalogue.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getType,
	}
	initGetTypeFlags(getTypeCommand.Flags())

	getTypeIssuesCommand := &cobra.Command{
		Use:                   `issues [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get type issues",
		Long:                  "The types in the catalogue have one or more issues for the different years, mintmarks and varieties.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getTypeIssues,
	}
	initGetTypeIssuesFlags(getTypeIssuesCommand.Flags())

	getTypeIssuesPricesCommand := &cobra.Command{
		Use:                   `prices [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get type issues prices",
		Long:                  "Get estimates for the price of an issue, depending on the grade.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getTypeIssuesPrices,
	}
	initGetTypeIssuesPricesFlags(getTypeIssuesPricesCommand.Flags())

	searchTypesCommand := &cobra.Command{
		Use:                   `search [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "search types",
		Long:                  "Search the catalogue for coin, banknote, and exonumia types.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  searchTypes,
	}
	initSearchTypesFlags(searchTypesCommand.Flags())

	mintsCommand := &cobra.Command{
		Use:                   `mints [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list mints",
		Long:                  "Retrieve the details about all the mints.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  mints,
	}
	initMintsFlags(mintsCommand.Flags())

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
		typesSubcommand,
		versionCommand,
	)

	typesSubcommand.AddCommand(
		getTypeCommand,
		getTypeIssuesCommand,
		getTypeIssuesPricesCommand,
		searchTypesCommand,
	)
	return rootCommand
}
