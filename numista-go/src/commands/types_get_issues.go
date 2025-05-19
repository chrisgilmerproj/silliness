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

func initGetTypeIssuesFlags(flag *pflag.FlagSet) {
	flag.Int32(flagTypesTypeID, 0, "ID of the type to fetch")
}

func validateGetTypeIssuesFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}

	typeID := v.GetInt32(flagTypesTypeID)
	if typeID < 0 {
		return fmt.Errorf("Type ID must be greater than or equal to 0")
	}
	return nil
}

func getTypeIssues(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateGetTypeIssuesFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	typeID := v.GetInt32(flagTypesTypeID)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiGetIssuesOpts{
		Lang: optional.NewString(lang),
	}

	issues, errSearchTypesID := numista.GetTypeIssues(apiClient, ctx, typeID, &opts)
	if errSearchTypesID != nil {
		return errSearchTypesID
	}

	jsonResponse, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}
	fmt.Println(string(jsonResponse))

	return nil
}
