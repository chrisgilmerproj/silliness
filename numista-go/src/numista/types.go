package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetType(apiClient *swagger.APIClient, ctx context.Context, typeID int32, opts *swagger.CatalogueApiGetTypeOpts) (*swagger.ModelType, error) {

	inlineResp, resp, errSearchTypes := apiClient.CatalogueApi.GetType(ctx, typeID, opts)
	if errSearchTypes != nil {
		return nil, errSearchTypes
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting type: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetTypeIssues(apiClient *swagger.APIClient, ctx context.Context, typeID int32, opts *swagger.CatalogueApiGetIssuesOpts) (*[]swagger.Issue, error) {

	inlineResp, resp, errSearchTypes := apiClient.CatalogueApi.GetIssues(ctx, typeID, opts)
	if errSearchTypes != nil {
		return nil, errSearchTypes
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting issues: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetTypeIssuesPrices(apiClient *swagger.APIClient, ctx context.Context, typeID, issueID int32, opts *swagger.CatalogueApiGetPricesOpts) (*swagger.InlineResponse2001, error) {

	inlineResp, resp, errSearchTypes := apiClient.CatalogueApi.GetPrices(ctx, typeID, issueID, opts)
	if errSearchTypes != nil {
		return nil, errSearchTypes
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting prices: %v", resp.Status)
	}
	return &inlineResp, nil
}

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
