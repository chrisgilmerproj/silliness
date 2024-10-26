package numista

import (
	"net/http"
	"os"
	"time"

	"github.com/chrisgilmerproj/silliness/numista-go/v2/src/swagger"
)

func NewAPIClient() *swagger.APIClient {
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

	return apiClient
}
