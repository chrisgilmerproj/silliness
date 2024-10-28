package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetCatalogues(apiClient *swagger.APIClient, ctx context.Context) (*swagger.InlineResponse2005, error) {
	inlineResp, resp, errGetCatalogues := apiClient.CatalogueApi.GetCatalogues(ctx)
	if errGetCatalogues != nil {
		return nil, errGetCatalogues
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting catalogues: %v", resp.Status)
	}
	return &inlineResp, nil
}
