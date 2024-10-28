package numista

import (
	"context"
	"fmt"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func GetUser(apiClient *swagger.APIClient, ctx context.Context, userID int32, opts *swagger.UserApiGetUserOpts) (*swagger.InlineResponse2008, error) {
	inlineResp, resp, errGetUser := apiClient.UserApi.GetUser(ctx, userID, opts)
	if errGetUser != nil {
		return nil, errGetUser
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting user: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetUserCollections(apiClient *swagger.APIClient, ctx context.Context, userID int32) (*swagger.InlineResponse2009, error) {
	inlineResp, resp, errGetUserCollections := apiClient.UserApi.GetUserCollections(ctx, userID)
	if errGetUserCollections != nil {
		return nil, errGetUserCollections
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting collections: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetUserCollectedItems(apiClient *swagger.APIClient, ctx context.Context, userID int32, opts *swagger.UserApiGetCollectedItemsOpts) (*swagger.InlineResponse20010, error) {
	inlineResp, resp, errGetCollectedItems := apiClient.UserApi.GetCollectedItems(ctx, userID, opts)
	if errGetCollectedItems != nil {
		return nil, errGetCollectedItems
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting collections: %v", resp.Status)
	}
	return &inlineResp, nil
}

func GetUserCollectedItem(apiClient *swagger.APIClient, ctx context.Context, userID, itemID int32) (*swagger.CollectedItem, error) {
	inlineResp, resp, errGetCollectedItems := apiClient.UserApi.GetCollectedItem(ctx, userID, itemID)
	if errGetCollectedItems != nil {
		return nil, errGetCollectedItems
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting collections: %v", resp.Status)
	}
	return &inlineResp, nil
}
