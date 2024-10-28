package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/text/currency"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

const (
	flagTypesCurrency = "currency"
)

func initGetTypeIssuesPricesFlags(flag *pflag.FlagSet) {
	flag.Int32(flagTypesTypeID, 0, "ID of the type to fetch")
	flag.Int32(flagTypesIssueID, 0, "ID of the issue to fetch")
	flag.String(flagTypesCurrency, currency.EUR.String(), "ID of the issue to fetch")
}

func validateGetTypeIssuesPricesFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}

	typeID := v.GetInt32(flagTypesTypeID)
	if typeID < 0 {
		return fmt.Errorf("Type ID must be greater than or equal to 0")
	}

	issueID := v.GetInt32(flagTypesIssueID)
	if issueID < 0 {
		return fmt.Errorf("Issue ID must be greater than or equal to 0")
	}

	cur := v.GetString(flagTypesCurrency)
	if cur == "" {
		return fmt.Errorf("Currency must be provided")
	}
	_, errParseCurrency := currency.ParseISO(cur)
	if errParseCurrency != nil {
		return fmt.Errorf("Currency must be a valid ISO 4217 currency code. You provided %s. Error: %w", cur, errParseCurrency)
	}
	return nil
}

func getTypeIssuesPrices(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidateRoot := validateRootFlags(v)
	if errValidateRoot != nil {
		return errValidateRoot
	}
	errValidate := validateGetTypeIssuesPricesFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	currency := v.GetString(flagTypesCurrency)
	typeID := v.GetInt32(flagTypesTypeID)
	issueID := v.GetInt32(flagTypesIssueID)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiGetPricesOpts{
		Lang:     optional.NewString(lang),
		Currency: optional.NewString(currency),
	}

	issues, errSearchTypesID := numista.GetTypeIssuesPrices(apiClient, ctx, typeID, issueID, &opts)
	if errSearchTypesID != nil {
		return errSearchTypesID
	}

	jsonResponse, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}
	fmt.Println(string(jsonResponse))

	return nil
}
