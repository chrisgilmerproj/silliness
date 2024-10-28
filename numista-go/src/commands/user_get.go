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
	flagUserUserID = "user-id"
)

func initGetUserFlags(flag *pflag.FlagSet) {
	flag.Int32(flagUserUserID, 0, "ID of the user to fetch")
}

func validateGetUserFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	userID := v.GetInt32(flagUserUserID)
	if userID < 1 {
		return fmt.Errorf("User ID must be greater than or equal to 1")
	}
	return nil
}

func getUser(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateGetUserFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	userID := v.GetInt32(flagUserUserID)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.UserApiGetUserOpts{
		Lang: optional.NewString(lang),
	}

	mints, errGetMint := numista.GetUser(apiClient, ctx, userID, &opts)
	if errGetMint != nil {
		return fmt.Errorf("error getting mint: %w", errGetMint)
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
