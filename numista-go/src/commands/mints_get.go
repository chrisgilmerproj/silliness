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

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

const (
	flagMintsMintID = "mint-id"
)

// Initialize the flags for the subcommand
func initGetMintFlags(flag *pflag.FlagSet) {
	flag.Int32(flagMintsMintID, 0, "ID of the mint to fetch")
}

func validateGetMintFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	mintID := v.GetInt32(flagMintsMintID)
	if mintID < 1 {
		return fmt.Errorf("Mint ID must be greater than or equal to 1")
	}
	return nil
}

func getMint(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateGetMintFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	mintID := v.GetInt32(flagMintsMintID)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiGetMintOpts{
		Lang: optional.NewString(lang),
	}

	mints, errGetMint := numista.GetMint(apiClient, ctx, mintID, &opts)
	if errGetMint != nil {
		return errGetMint
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
