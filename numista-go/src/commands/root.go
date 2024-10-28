package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	API_KEY_NAME = "NUMISTA_API_KEY"

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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v, errViper := initViper(cmd)
			if errViper != nil {
				return fmt.Errorf("error initializing viper: %w", errViper)
			}

			// Validate the api key exists by checking environment variables
			_, exists := os.LookupEnv(API_KEY_NAME)
			if !exists {
				return fmt.Errorf("API key must be set in the environment variable %s", API_KEY_NAME)
			}

			// Validate the root flags
			errValidate := validateRootFlags(v)
			if errValidate != nil {
				return errValidate
			}
			return nil
		},
	}
	initRootFlags(rootCommand.PersistentFlags())

	cataloguesSubcommand := &cobra.Command{
		Use:                   `catalogues [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "catalogues commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	listCataloguesCommand := &cobra.Command{
		Use:                   `list [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list catalogues",
		Long:                  "Retrieve the list of all the reference catalogues used for cross-reference in the catalogue.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  listCatalogues,
	}
	initListCataloguesFlags(listCataloguesCommand.Flags())

	issuersSubcommand := &cobra.Command{
		Use:                   `issuers [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "issuers commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	listIssuersCommand := &cobra.Command{
		Use:                   `list [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list issuers",
		Long:                  "Retrieve the details about all the issuing countries and territories.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  listIssuers,
	}
	initListIssuersFlags(listIssuersCommand.Flags())

	mintsSubcommand := &cobra.Command{
		Use:                   `mints [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "mints commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	listMintsCommand := &cobra.Command{
		Use:                   `list [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list mints",
		Long:                  "Retrieve the details about all the mints.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  listMints,
	}
	initListMintsFlags(listMintsCommand.Flags())

	getMintCommand := &cobra.Command{
		Use:                   `get [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get mint",
		Long:                  "Retrieve the details about a specific mint.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getMint,
	}
	initGetMintFlags(getMintCommand.Flags())

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

	userSubcommand := &cobra.Command{
		Use:                   `user [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "user commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	getUserCommand := &cobra.Command{
		Use:                   `get [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get user",
		Long:                  "Retrieve the details about a specific user.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getUser,
	}
	initGetUserFlags(getUserCommand.Flags())

	getUserCollectionsCommand := &cobra.Command{
		Use:                   `collections [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get collections",
		Long:                  "A user can organize their items in multiple collections.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getUserCollections,
	}
	initGetUserCollectionsFlags(getUserCollectionsCommand.Flags())

	collectedItemsSubcommand := &cobra.Command{
		Use:                   `collected-items [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "collected items commands",
		SilenceErrors:         true,
		SilenceUsage:          true,
	}

	getUserCollectedItemsCommand := &cobra.Command{
		Use:                   `list [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "list collected items",
		Long:                  "Retrieve the details about all the items (coins, banknotes, pieces of exonumia) in collected by the user.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getUserCollectedItems,
	}
	initGetUserCollectedItemsFlags(getUserCollectedItemsCommand.Flags())

	getUserCollectedItemCommand := &cobra.Command{
		Use:                   `get [flags]`,
		DisableFlagsInUseLine: true,
		Short:                 "get collected item",
		Long:                  "Retrieve the details about a specific item collected by the user.",
		SilenceErrors:         true,
		SilenceUsage:          true,
		RunE:                  getUserCollectedItem,
	}
	initGetUserCollectedItemFlags(getUserCollectedItemCommand.Flags())

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
		cataloguesSubcommand,
		issuersSubcommand,
		mintsSubcommand,
		typesSubcommand,
		userSubcommand,
		versionCommand,
	)

	cataloguesSubcommand.AddCommand(
		listCataloguesCommand,
	)

	issuersSubcommand.AddCommand(
		listIssuersCommand,
	)

	mintsSubcommand.AddCommand(
		listMintsCommand,
		getMintCommand,
	)

	typesSubcommand.AddCommand(
		getTypeCommand,
		getTypeIssuesCommand,
		getTypeIssuesPricesCommand,
		searchTypesCommand,
	)

	collectedItemsSubcommand.AddCommand(
		getUserCollectedItemsCommand,
		getUserCollectedItemCommand,
	)

	userSubcommand.AddCommand(
		getUserCommand,
		getUserCollectionsCommand,
		collectedItemsSubcommand,
	)

	return rootCommand
}
