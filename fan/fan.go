package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	owm "github.com/briandowns/openweathermap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	smartdevicemanagement "google.golang.org/api/smartdevicemanagement/v1"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func authenticate(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	log.Print("Your browser will be opened to authenticate with Google")
	log.Print("Hit enter to confirm and continue")

	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	codeChannel := make(chan string)
	mux := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: mux}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p style=font-size:xx-large;text-align:center>return to your terminal</p>"))

		codeChannel <- r.URL.Query().Get("code")
	})

	openBrowser(url)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	code := <-codeChannel

	log.Print("Shutting down server")

	server.Shutdown(context.Background())

	log.Print("Exchanging token")

	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

type TraitsConnectivity struct {
	Status string `json:"status"`
}
type TraitsFan struct {
	TimerMode string `json:"timerMode"`
}
type TraitsHumidity struct {
	AmbientHumidityPercent int `json:"ambientHumidityPercent"`
}
type TraitsInfo struct {
	CustomName *string `json:"customName"`
}
type TraitsSettings struct {
	TemperatureScale string `json:"temperatureScale"`
}
type TraitsTemperature struct {
	AmbientTemperatureCelsius float64 `json:"ambientTemperatureCelsius"`
}
type TraitsThermostatEco struct {
	AvailableModes []string `json:"availableModes"`
	CoolCelsius    float64  `json:"coolCelsius"`
	HeatCelsius    float64  `json:"heatCelsius"`
	Mode           string   `json:"mode"`
}
type TraitsThermostatHvac struct {
	Status string `json:"status"`
}
type TraitsThermostatMode struct {
	AvailableModes []string `json:"availableModes"`
	Mode           string   `json:"mode"`
}
type TraitsThermostatTemperatureSetpoint struct {
	CoolCelsius float64 `json:"coolCelsius"`
}

type Traits struct {
	Connectivity                  TraitsConnectivity                  `json:"sdm.devices.traits.Connectivity"`
	Fan                           TraitsFan                           `json:"sdm.devices.traits.Fan"`
	Humidity                      TraitsHumidity                      `json:"sdm.devices.traits.Humidity"`
	Info                          TraitsInfo                          `json:"sdm.devices.traits.Info"`
	Settings                      TraitsSettings                      `json:"sdm.devices.traits.Settings"`
	Temperature                   TraitsTemperature                   `json:"sdm.devices.traits.Temperature"`
	ThermostatEco                 TraitsThermostatEco                 `json:"sdm.devices.traits.ThermostatEco"`
	ThermostatHvac                TraitsThermostatHvac                `json:"sdm.devices.traits.ThermostatHvac"`
	ThermostatMode                TraitsThermostatMode                `json:"sdm.devices.traits.ThermostatMode"`
	ThermostatTemperatureSetpoint TraitsThermostatTemperatureSetpoint `json:"sdm.devices.traits.ThermostatTemperatureSetpoint"`
}

func celsiusToFarenheit(temp float64) float64 {
	return temp*(9/5) + 32.0
}

// https://blog.mattcanty.com/2020-11-17-accessing-nest-thermostat-with-go/
func main() {

	// Get secrets from the environment
	clientID := os.Getenv("GCP_CLIENT_ID")
	clientSecret := os.Getenv("GCP_CLIENT_SECRET")
	projectID := os.Getenv("GCP_PROJECT_ID")
	OwmApiKey := os.Getenv("OWM_API_KEY")
	tmpTokenFile := "/tmp/gcpOauthToken"

	ctx := context.Background()
	var token *oauth2.Token
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080",
		Scopes: []string{
			smartdevicemanagement.SdmServiceScope,
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://nestservices.google.com/partnerconnections/%s/auth", projectID),
			TokenURL: google.Endpoint.TokenURL,
		},
	}

	if _, err := os.Stat(tmpTokenFile); !errors.Is(err, os.ErrNotExist) {
		bToken, errReadFile := os.ReadFile(tmpTokenFile)
		if errReadFile != nil {
			fmt.Println(errReadFile)
			return
		}
		errUnmarshal := json.Unmarshal(bToken, &token)
		if errUnmarshal != nil {
			fmt.Println(errUnmarshal)
			return
		}
	}

	if token == nil || (token != nil && token.Expiry.Before(time.Now())) {
		token = authenticate(ctx, conf)
		bToken, errMarshal := json.Marshal(token)
		if errMarshal != nil {
			fmt.Println(errMarshal)
			return
		}
		errWriteFile := os.WriteFile(tmpTokenFile, bToken, 644)
		if errWriteFile != nil {
			fmt.Println(errWriteFile)
			return
		}
	}

	httpClient := conf.Client(ctx, token)
	sdmService, err := smartdevicemanagement.NewService(ctx,
		option.WithScopes(smartdevicemanagement.SdmThermostatServiceScope),
		option.WithHTTPClient(httpClient),
	)
	if err != nil {
		fmt.Println(err)
	}

	w, errNewCurrent := owm.NewCurrent("F", "EN", OwmApiKey)
	if errNewCurrent != nil {
		log.Fatalln(errNewCurrent)
	}

	// Continue to get the data
	for {
		devicesListResponse, err := sdmService.Enterprises.Devices.List(fmt.Sprintf("enterprises/%s", projectID)).Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, device := range devicesListResponse.Devices {

			var nestTraits Traits
			err = json.Unmarshal(device.Traits, &nestTraits)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Inside  Temperature: %.2f°F", celsiusToFarenheit(nestTraits.Temperature.AmbientTemperatureCelsius))
			log.Printf("Inside  Humidity:    %d%%", nestTraits.Humidity.AmbientHumidityPercent)
		}

		w.CurrentByZip(90807, "US")
		log.Printf("Outside Temperature: %.2F°F", w.Main.Temp)
		log.Printf("Outside Humidity:    %d%%", w.Main.Humidity)
		time.Sleep(5 * time.Second)
	}
}
