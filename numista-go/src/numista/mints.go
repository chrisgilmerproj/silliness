package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetMints(apiClient *swagger.APIClient, ctx context.Context, opts *swagger.CatalogueApiGetMintsOpts) (*swagger.InlineResponse2004, error) {
	inlineResp, resp, errGetMints := apiClient.CatalogueApi.GetMints(ctx, opts)
	if errGetMints != nil {
		return nil, errGetMints
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting mints: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetMint(apiClient *swagger.APIClient, ctx context.Context, mintID int32, opts *swagger.CatalogueApiGetMintOpts) (*swagger.Mint, error) {

	inlineResp, resp, errSearchMint := apiClient.CatalogueApi.GetMint(ctx, mintID, opts)
	if errSearchMint != nil {
		return nil, errSearchMint
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting mint: %v", resp.Status)
	}
	return &inlineResp, nil
}
