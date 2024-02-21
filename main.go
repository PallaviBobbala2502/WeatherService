package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ResponseEvent
type WheatherResponseEvent struct {
	Temperature      string
	WeatherCondition string
	Location         string
}

func main() {
	//create new router using gorilla muxs framework
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/weather", WeatherHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("API is up and running...")

}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat, lon, apiKey := getInputParams(r)
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

func getInputParams(r *http.Request) (string, string, string) {
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("lon")
	apiKey := r.URL.Query().Get("apikey")

	return lat, long, apiKey
}

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

func getTemperatureCondition(temperature float64) string {
	temperatureCelsius := temperature - 273.15
	temperatureCondition := "moderate"
	if temperatureCelsius < 10 {
		temperatureCondition = "cold"
	} else if temperatureCelsius > 30 {
		temperatureCondition = "hot"
	}
	return temperatureCondition
}
