package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func main() {

	apiClient := swagger.NewAPIClient(&swagger.Configuration{
		BasePath: "https://api.numista.com/v3",
		DefaultHeader: map[string]string{
			"Numista-API-Key": os.Getenv("NUMISTA_API_KEY"),
		},
		UserAgent: "Swagger-Codegen/1.0.0/go",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	})

	ctx := context.Background()

	lang := optional.NewString("en")
	inlineResp, resp, errGetMints := apiClient.CatalogueApi.GetMints(ctx, &swagger.CatalogueApiGetMintsOpts{Lang: lang})
	if errGetMints != nil {
		log.Fatalf("Error getting mints: %v", errGetMints)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Error getting mints: %v", resp.Status)
	}

	for _, mint := range inlineResp.Mints {
		fmt.Printf("Mint: %s\n", mint.Name)
	}

	// Marshal the response to JSON
	jsonResponse, err := json.MarshalIndent(inlineResp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting response as JSON: %v", err)
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))
}
