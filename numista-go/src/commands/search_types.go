package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/numista"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func searchTypes(cmd *cobra.Command) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}
	lang := v.GetString(flagLang)
	category := v.GetString("category")
	q := v.GetString("q")
	issuer := v.GetString("issuer")
	catalogue := v.GetInt32("catalogue")
	number := v.GetString("number")
	ruler := v.GetInt32("ruler")
	material := v.GetInt32("material")
	year := v.GetString("year")
	date := v.GetString("date")
	size := v.GetString("size")
	weight := v.GetString("weight")
	page := v.GetInt32("page")
	count := v.GetInt32("count")

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	// If q, issuer, catalogue, date, and year are all nil, return an error
	if q == "" && issuer == "" && catalogue == 0 && date == "" && year == "" {
		return fmt.Errorf("At least one of q, issuer, catalogue, date, or year must be provided")
	}
	opts := swagger.CatalogueApiSearchTypesOpts{
		Lang:      optional.NewString(lang),
		Category:  optional.NewString(category),
		Q:         optional.NewString(q),
		Issuer:    optional.NewString(issuer),
		Catalogue: optional.NewInt32(catalogue),
		Number:    optional.NewString(number),
		Ruler:     optional.NewInt32(ruler),
		Material:  optional.NewInt32(material),
		Year:      optional.NewString(year),
		Date:      optional.NewString(date),
		Size:      optional.NewString(size),
		Weight:    optional.NewString(weight),
		Page:      optional.NewInt32(page),
		Count:     optional.NewInt32(count),
	}

	types, errSearchTypes := numista.SearchTypes(apiClient, ctx, &opts)
	if errSearchTypes != nil {
		return errSearchTypes
	}

	jsonResponse, err := json.MarshalIndent(types, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}
	fmt.Println(string(jsonResponse))

	return nil
}
