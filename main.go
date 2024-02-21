package main

import (
	"log"
	"net/http"

	handler "github.com/PallaviBobbala2502/WeatherService/weatherHandler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Initiating weather api...")
	err := godotenv.Load("config/local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//create new router using gorilla muxs framework
	router := mux.NewRouter()

	//specify endpoints, handler functions
	router.HandleFunc("/weather", handler.WeatherHandler).Methods("GET")
	log.Println("API is up and running...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
