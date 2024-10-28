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
	flagTypesCategory   = "category"
	flagTypesQuery      = "query"
	flagTypesQueryShort = "q"
	flagTypesIssuer     = "issuer"
	flagTypesCatalogue  = "catalogue"
	flagTypesNumber     = "number"
	flagTypesRuler      = "ruler"
	flagTypesMaterial   = "material"
	flagTypesYear       = "year"
	flagTypesDate       = "date"
	flagTypesSize       = "size"
	flagTypesWeight     = "weight"
	flagTypesPage       = "page"
	flagTypesCount      = "count"
)

var (
	typesCategoryChoices = []string{
		"coin",
		"banknote",
		"exonumia",
	}
)

func initSearchTypesFlags(flag *pflag.FlagSet) {
	flag.String(flagTypesCategory, "", "Catalogue category")
	flag.StringP(flagTypesQuery, flagTypesQueryShort, "", "Search query")
	flag.String(flagTypesIssuer, "", "Issuer code. If provided, only the types from the given issuer are returned.")
	flag.Int32(flagTypesCatalogue, 0, "ID of a reference catalogue. If provided, only the types referenced in the given catalogue are returned.")
	flag.String(flagTypesNumber, "", "Number of the searched typed in a reference catalogue. This parameter works only with the other parameter catalogue.")
	flag.Int32(flagTypesRuler, 0, "ID of a ruler. If provided, only the types issued by this ruler are returned.")
	flag.Int32(flagTypesMaterial, 0, "ID of a material. If provided, only the types made of this material are returned.")
	flag.String(flagTypesYear, "", "Year, or a range of years, as given on the type")
	flag.String(flagTypesDate, "", "Gregorian date when the type was issued, or a range of dates.")
	flag.String(flagTypesSize, "", "Main (largest) dimension of the type specimen, or boundaries of dimensions. In the first case, the search allows a ±0.5 millimeters tolerance.")
	flag.String(flagTypesWeight, "", "Standard or average weight of the type, or boundaries of weight. In the first case, the search allows a ±0.5 grams tolerance.")
	flag.Int32(flagTypesPage, 1, "Page of results")
	flag.Int32(flagTypesCount, 50, "Results per page")
}

func validateSearchTypeFlags(v *viper.Viper, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("no positional arguments allowed")
	}

	category := v.GetString(flagTypesCategory)
	query := v.GetString(flagTypesQuery)
	issuer := v.GetString(flagTypesIssuer)
	catalogue := v.GetInt32(flagTypesCatalogue)
	// number := v.GetString(flagTypesNumber)
	ruler := v.GetInt32(flagTypesRuler)
	material := v.GetInt32(flagTypesMaterial)
	year := v.GetString(flagTypesYear)
	date := v.GetString(flagTypesDate)
	// size := v.GetString(flagTypesSize)
	// weight := v.GetString(flagTypesWeight)
	page := v.GetInt32(flagTypesPage)
	count := v.GetInt32(flagTypesCount)

	// If q, issuer, catalogue, date, and year are all nil, return an error
	if query == "" && issuer == "" && catalogue == 0 && date == "" && year == "" {
		return fmt.Errorf("At least one of q, issuer, catalogue, date, or year must be provided")
	}

	if category != "" {
		if !utils.Contains(typesCategoryChoices, category) {
			return fmt.Errorf("category must be one of %v", typesCategoryChoices)
		}
	}

	if catalogue < 0 {
		return fmt.Errorf("catalogue must be greater than or equal to 0")
	}
	if ruler < 0 {
		return fmt.Errorf("ruler must be greater than or equal to 0")
	}
	if material < 0 {
		return fmt.Errorf("material must be greater than or equal to 0")
	}
	if page < 1 {
		return fmt.Errorf("page must be greater than or equal to 1")
	}
	if count < 1 || count > 50 {
		return fmt.Errorf("count must be greater than or equal to 1 and less than or equal to 50")
	}

	return nil
}

func searchTypes(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	errValidate := validateSearchTypeFlags(v, args)
	if errValidate != nil {
		return errValidate
	}

	lang := v.GetString(flagLang)
	category := v.GetString(flagTypesCategory)
	query := v.GetString(flagTypesQuery)
	issuer := v.GetString(flagTypesIssuer)
	catalogue := v.GetInt32(flagTypesCatalogue)
	number := v.GetString(flagTypesNumber)
	ruler := v.GetInt32(flagTypesRuler)
	material := v.GetInt32(flagTypesMaterial)
	year := v.GetString(flagTypesYear)
	date := v.GetString(flagTypesDate)
	size := v.GetString(flagTypesSize)
	weight := v.GetString(flagTypesWeight)
	page := v.GetInt32(flagTypesPage)
	count := v.GetInt32(flagTypesCount)

	ctx := context.Background()

	apiClient := numista.NewAPIClient()

	opts := swagger.CatalogueApiSearchTypesOpts{
		Lang:  optional.NewString(lang),
		Page:  optional.NewInt32(page),
		Count: optional.NewInt32(count),
	}
	if category != "" {
		opts.Category = optional.NewString(category)
	}
	if query != "" {
		opts.Q = optional.NewString(query)
	}
	if issuer != "" {
		opts.Issuer = optional.NewString(issuer)
	}
	if catalogue != 0 {
		opts.Catalogue = optional.NewInt32(catalogue)
	}
	if number != "" {
		opts.Number = optional.NewString(number)
	}
	if ruler != 0 {
		opts.Ruler = optional.NewInt32(ruler)
	}
	if material != 0 {
		opts.Material = optional.NewInt32(material)
	}
	if year != "" {
		opts.Year = optional.NewString(year)
	}
	if date != "" {
		opts.Date = optional.NewString(date)
	}
	if size != "" {
		opts.Size = optional.NewString(size)
	}
	if weight != "" {
		opts.Weight = optional.NewString(weight)
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
