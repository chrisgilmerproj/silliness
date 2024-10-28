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
	flagTypesTypeID  = "type-id"
	flagTypesIssueID = "issue-id"
)

func initGetTypeFlags(flag *pflag.FlagSet) {
	flag.Int32(flagTypesTypeID, 0, "ID of the type to fetch")
}

func validateGetTypeFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}

	typeID := v.GetInt32(flagTypesTypeID)
	if typeID < 0 {
		return fmt.Errorf("Type ID must be greater than or equal to 0")
	}
	return nil
}

func getType(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidateRoot := validateRootFlags(v)
	if errValidateRoot != nil {
		return errValidateRoot
	}
	errValidate := validateGetTypeFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	typeID := v.GetInt32(flagTypesTypeID)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiGetTypeOpts{
		Lang: optional.NewString(lang),
	}

	types, errSearchTypesID := numista.GetType(apiClient, ctx, typeID, &opts)
	if errSearchTypesID != nil {
		return errSearchTypesID
	}

	jsonResponse, err := json.MarshalIndent(types, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}
	fmt.Println(string(jsonResponse))

	return nil
}
