package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func initListIssuersFlags(flag *pflag.FlagSet) {
}

func validateListIssuersFlags(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	return nil
}

func listIssuers(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateListIssuersFlags(args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiGetIssuersOpts{Lang: optional.NewString(lang)}

	issuers, errGetIssuers := numista.GetIssuers(apiClient, ctx, &opts)
	if errGetIssuers != nil {
		return errGetIssuers
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(issuers, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
