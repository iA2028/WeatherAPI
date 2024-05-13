package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/weather", Handler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract latitude and longitude from the query string
	latitude := r.URL.Query().Get("lat")
	longitude := r.URL.Query().Get("lon")

	// Check if latitude and longitude are provided
	if latitude == "" || longitude == "" {
		http.Error(w, "Missing latitude or longitude parameter", http.StatusBadRequest)
		return
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		http.Error(w, "Missing API key", http.StatusInternalServerError)
		return
	}

	// Make a request to the OpenWeather API
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", latitude, longitude, apiKey)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Parse the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
		return
	}

	// Extract weather information
	weather := data["weather"].([]interface{})
	if len(weather) == 0 {
		http.Error(w, "Weather information not found", http.StatusNotFound)
		return
	}
	weatherCondition := weather[0].(map[string]interface{})["main"].(string)
	temperature := data["main"].(map[string]interface{})["temp"].(float64)

	// Determine temperature description
	var tempDescription string
	switch {
	case temperature < 10:
		tempDescription = "cold"
	case temperature >= 10 && temperature < 20:
		tempDescription = "moderate"
	default:
		tempDescription = "hot"
	}

	// Create the weatherResponse
	weatherResponse := struct {
		WeatherCondition string  `json:"weatherCondition"`
		Temperature      float64 `json:"temperature"`
		TempDescription  string  `json:"tempDescription"`
	}{
		WeatherCondition: weatherCondition,
		Temperature:      temperature,
		TempDescription:  tempDescription,
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response as JSON
	json.NewEncoder(w).Encode(weatherResponse)
}
