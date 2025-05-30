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

func initGetUserCollectionsFlags(flag *pflag.FlagSet) {
	flag.Int32(flagUserUserID, 0, "ID of the user to fetch")
}

func validateGetUserCollectionsFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	return nil
}

func getUserCollections(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateGetUserCollectionsFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	oauthToken, userId, errGetOAuthToken := numista.Auth(apiClient, ctx)
	if errGetOAuthToken != nil {
		return errGetOAuthToken
	}

	apiClient = numista.NewOAuthClient(oauthToken)

	collections, errGetCollections := numista.GetUserCollections(apiClient, ctx, userId)
	if errGetCollections != nil {
		return errGetCollections
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(collections, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
