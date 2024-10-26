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

// Initialize the flags for the subcommand
func initMintsFlags(flag *pflag.FlagSet) {
}

func mints(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	lang := v.GetString(flagLang)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	mints, errGetMints := numista.GetMints(apiClient, ctx, optional.NewString(lang))
	if errGetMints != nil {
		return fmt.Errorf("error getting mints: %w", errGetMints)
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
