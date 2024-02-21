# WeatherService
WeatherService api-  to get current weather status based on the latitude and longitude coordinates.

# Clone the project
 $ git clone https://github.com/PallaviBobbala2502/WeatherService.git
 $ cd WeatherService

 # Run 
 $ go run main.go

# Environment Variables
 Load the APIKEY from env

 # /weather API- GET
 # Input Params
  Pass the input params lat & lon 
  Example: Test API in local (Postman)
  URL: localhost:8080/weather?lat=39.7392&lon=104.9903

# /weather API - Output
 Example:
 {
    "Temperature": "cold",
    "WeatherCondition": "Clouds",
    "Location": "Bayan Hot"
 }

