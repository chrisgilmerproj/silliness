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
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/utils"
)

const (
	flagUserCollectedItemsCategory   = "category"
	flagUserCollectedItemsType       = "type"
	flagUserCollectedItemsCollection = "collection"
)

var (
	UserCollectedItemsCategory = []string{"coin", "banknote", "exonumia"}
)

func initGetUserCollectedItemsFlags(flag *pflag.FlagSet) {
	flag.String(flagUserCollectedItemsCategory, "", "Category of the item")
	flag.Int32(flagUserCollectedItemsType, 0, "Type of the item")
	flag.Int32(flagUserCollectedItemsCollection, 0, "Collection of the item")
}

func validateGetUserCollectedItemsFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}
	category := v.GetString(flagUserCollectedItemsCategory)
	if category != "" {
		if !utils.Contains(UserCollectedItemsCategory, category) {
			return fmt.Errorf("category must be one of %v", UserCollectedItemsCategory)
		}
	}

	itemType := v.GetInt32(flagUserCollectedItemsType)
	if itemType < 0 {
		return fmt.Errorf("Type must be greater than or equal to 0")
	}
	collection := v.GetInt32(flagUserCollectedItemsCollection)
	if collection < 0 {
		return fmt.Errorf("Collection must be greater than or equal to 0")
	}

	return nil
}

func getUserCollectedItems(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateGetUserCollectedItemsFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	category := v.GetString(flagUserCollectedItemsCategory)
	itemType := v.GetInt32(flagUserCollectedItemsType)
	collection := v.GetInt32(flagUserCollectedItemsCollection)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	oauthToken, userId, errGetOAuthToken := numista.Auth(apiClient, ctx)
	if errGetOAuthToken != nil {
		return errGetOAuthToken
	}

	apiClient = numista.NewOAuthClient(oauthToken)

	opts := &swagger.UserApiGetCollectedItemsOpts{}
	if category != "" {
		opts.Category = optional.NewString(category)
	}
	if itemType > 0 {
		opts.Type_ = optional.NewInt32(itemType)
	}
	if collection > 0 {
		opts.Collection = optional.NewInt32(collection)
	}

	collectedItems, errGetCollectedItems := numista.GetUserCollectedItems(apiClient, ctx, userId, opts)
	if errGetCollectedItems != nil {
		return errGetCollectedItems
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(collectedItems, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))

	return nil
}
