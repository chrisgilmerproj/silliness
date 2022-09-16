package main

import (
	"fmt"
	"log"
	"os"
	"time"

	owm "github.com/briandowns/openweathermap"
)

func main() {

	// Get secrets from the environment
	OwmApiKey := os.Getenv("OWM_API_KEY")

	// Continue to get the data
	for {

		w, err := owm.NewCurrent("F", "EN", OwmApiKey) // fahrenheit (imperial) with Russian output
		// w, err := owm.NewOneCall("F", "EN", OwmApiKey, []string{})
		if err != nil {
			log.Fatalln(err)
		}

		w.CurrentByZip(90807, "US")
		/*
		 w.OneCallByCoordinates(
		 	&owm.Coordinates{
		 		Longitude: -118.171690,
		 		Latitude:  33.828440,
		 	},
		 )
		*/
		//fmt.Printf("%+v\n", w)
		fmt.Println(w.Main.Temp)
		time.Sleep(5 * time.Second)
	}
}
