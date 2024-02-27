package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/natansa/temperatura-cep/services"
)

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		zipcode := queryParams.Get("zipcode")

		if len(zipcode) != 8 {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		zipcodeHandler := services.NewZipcodeHandler()
		weatherService := services.NewWeatherService()

		cityName, err := zipcodeHandler.FetchCityNameFromZipcode(zipcode)
		if err != nil {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		tempCelsius, err := weatherService.FetchWeather(cityName)
		if err != nil {
			http.Error(w, "error fetching weather information", http.StatusInternalServerError)
			return
		}

		tempFahrenheit := services.CelsiusToFahrenheit(tempCelsius)
		tempKelvin := services.CelsiusToKelvin(tempCelsius)

		response := map[string]float64{
			"temp_C": tempCelsius,
			"temp_F": tempFahrenheit,
			"temp_K": tempKelvin,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
