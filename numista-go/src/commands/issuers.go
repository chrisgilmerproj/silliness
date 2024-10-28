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
)

func initIssuersFlags(flag *pflag.FlagSet) {
}

func issuers(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidateRoot := validateRootFlags(v)
	if errValidateRoot != nil {
		return errValidateRoot
	}

	lang := v.GetString(flagLang)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	mints, errGetIssuers := numista.GetIssuers(apiClient, ctx, optional.NewString(lang))
	if errGetIssuers != nil {
		return fmt.Errorf("error getting mints: %w", errGetIssuers)
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(mints, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
