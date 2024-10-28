package numista

import (
	"context"
	"fmt"

	"github.com/antihax/optional"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetMints(apiClient *swagger.APIClient, ctx context.Context, lang optional.String) (*swagger.InlineResponse2004, error) {
	inlineResp, resp, errGetMints := apiClient.CatalogueApi.GetMints(ctx, &swagger.CatalogueApiGetMintsOpts{Lang: lang})
	if errGetMints != nil {
		return nil, errGetMints
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting mints: %v", resp.Status)
	}
	return &inlineResp, nil
}
