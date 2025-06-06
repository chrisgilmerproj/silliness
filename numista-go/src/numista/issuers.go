package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetIssuers(apiClient *swagger.APIClient, ctx context.Context, opts *swagger.CatalogueApiGetIssuersOpts) (*swagger.InlineResponse2003, error) {
	inlineResp, resp, errGetIssuers := apiClient.CatalogueApi.GetIssuers(ctx, opts)
	if errGetIssuers != nil {
		return nil, errGetIssuers
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting issuers: %v", resp.Status)
	}
	return &inlineResp, nil
}
