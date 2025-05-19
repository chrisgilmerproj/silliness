package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
)

func authUser(cmd *cobra.Command, args []string) error {
	_, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	response, _, errGetOauthToken := numista.Auth(apiClient, ctx)
	if errGetOauthToken != nil {
		return errGetOauthToken
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
