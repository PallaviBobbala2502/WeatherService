package weatherHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

//const APIKEY = "3c3550e1a10b9c6306f3f46493c3f5ec"

// ResponseEvent
type WheatherResponseEvent struct {
	Temperature      string
	WeatherCondition string
	Location         string
}

// WeatherHandler
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat, lon, err := getInputParams(r)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	apiKey := os.Getenv("API_KEY")

	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&apiKey=%v", lat, lon, apiKey)

	resp, err := http.Get(URL)
	if err != nil {
		http.Error(w, "error occured while reading request body", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error occured while reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(res, &data); err != nil {
		http.Error(w, "error occured while unmarshalling the data", http.StatusInternalServerError)
		return
	}

	response := getResponse(data)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error occured while marshalling the data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

// getInputParams - Return latitude and longitude coordinates
func getInputParams(r *http.Request) (float64, float64, error) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		return 0, 0, err
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		return 0, 0, err
	}
	//apiKey := r.URL.Query().Get("apikey")

	return lat, lon, nil
}

// getResponse
func getResponse(data map[string]interface{}) WheatherResponseEvent {
	var response WheatherResponseEvent
	if data["weather"] != nil {
		weather := data["weather"].([]interface{})[0].(map[string]interface{})["main"].(string)
		temperature := data["main"].(map[string]interface{})["temp"].(float64)
		location := data["name"].(string)
		response = WheatherResponseEvent{
			WeatherCondition: weather,
			Temperature:      getTemperatureCondition(temperature),
			Location:         location,
		}
	}
	return response
}

// getTemperatureCondition - Return temperature condition based on temperature in celsius
func getTemperatureCondition(temperature float64) string {
	temperatureInCelsius := temperature - 273.15
	temperatureCondition := "moderate"
	if temperatureInCelsius < 10 {
		temperatureCondition = "cold"
	} else if temperatureInCelsius > 20 {
		temperatureCondition = "hot"
	}
	return temperatureCondition
}
