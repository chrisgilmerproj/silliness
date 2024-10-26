package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func SearchTypes(apiClient *swagger.APIClient, ctx context.Context, opts *swagger.CatalogueApiSearchTypesOpts) (*swagger.InlineResponse200, error) {

	inlineResp, resp, errSearchTypes := apiClient.CatalogueApi.SearchTypes(ctx, opts)
	if errSearchTypes != nil {
		return nil, errSearchTypes
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting mints: %v", resp.Status)
	}
	return &inlineResp, nil
}
