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
	inlineResp, resp, errGetUser := apiClient.UserApi.GetUserCollections(ctx, userID)
	if errGetUser != nil {
		return nil, errGetUser
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting collections: %v", resp.Status)
	}
	return &inlineResp, nil
}
