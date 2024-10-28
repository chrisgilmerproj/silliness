package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
)

func initListCataloguesFlags(flag *pflag.FlagSet) {
}

func validateListCataloguesFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	return nil
}

func listCatalogues(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateListCataloguesFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	catalogues, errGetCatalogues := numista.GetCatalogues(apiClient, ctx)
	if errGetCatalogues != nil {
		return errGetCatalogues
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(catalogues, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
