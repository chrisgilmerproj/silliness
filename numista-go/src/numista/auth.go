package numista

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

const tokenFilePath = "/tmp/numista_oauth_token.json"

type TokenData struct {
	AccessToken  string `json:"access_token"`
	ExpiresAt    int64  `json:"expires_at"`
	TokenType    string `json:"token_type"`
	ExpiresInSec int32  `json:"expires_in"`
	UserId       int32  `json:"user_id"`
}

func NewToken(inlineResp *swagger.InlineResponse2007) *TokenData {
	return &TokenData{
		AccessToken:  inlineResp.AccessToken,
		ExpiresAt:    time.Now().Unix() + int64(inlineResp.ExpiresIn),
		TokenType:    inlineResp.TokenType,
		ExpiresInSec: inlineResp.ExpiresIn,
		UserId:       inlineResp.UserId,
	}
}

func GetOAuthToken(apiClient *swagger.APIClient, ctx context.Context) (*swagger.InlineResponse2007, error) {
	opts := &swagger.OAuthApiGetOAuthTokenOpts{
		GrantType: optional.NewString("client_credentials"),
		Scope:     optional.NewString("view_collection"),
	}
	inlineResp, resp, errGetOAuthToken := apiClient.OAuthApi.GetOAuthToken(ctx, opts)
	if errGetOAuthToken != nil {
		return nil, errGetOAuthToken
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting OAuth token: %v", resp.Status)
	}
	return &inlineResp, nil
}

// Auth retrieves or refreshes the OAuth token
func Auth(apiClient *swagger.APIClient, ctx context.Context) (string, int32, error) {
	// Step 1: Check if the token file exists and is valid
	if tokenExistsAndValid() {
		token, err := loadTokenFromFile()
		if err == nil {
			return token.AccessToken, token.UserId, nil
		}
	}

	// Step 2: If the token is not valid or file doesn't exist, get a new token
	tokenData, err := GetOAuthToken(apiClient, ctx)
	if err != nil {
		return "", 0, err
	}

	// Save the token and expiration time in the file
	err = saveTokenToFile(tokenData)
	if err != nil {
		return "", 0, fmt.Errorf("failed to save token to file: %v", err)
	}

	return tokenData.AccessToken, tokenData.UserId, nil
}

func tokenExistsAndValid() bool {
	// Check if the file exists
	if _, err := os.Stat(tokenFilePath); os.IsNotExist(err) {
		return false
	}

	// Load the token and check if it has expired
	token, err := loadTokenFromFile()
	if err != nil {
		return false
	}

	// Check if the token is still valid (current time < expiration time)
	currentTime := time.Now().Unix()
	if currentTime < token.ExpiresAt {
		return true
	}

	return false
}

func loadTokenFromFile() (*TokenData, error) {
	data, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read token file: %v", err)
	}

	var token TokenData
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to parse token file: %v", err)
	}

	return &token, nil
}

func saveTokenToFile(inlineResp *swagger.InlineResponse2007) error {
	token := NewToken(inlineResp)

	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token data: %v", err)
	}

	return os.WriteFile(tokenFilePath, data, 0644)
}
