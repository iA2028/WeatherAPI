OpenWeather API Integration
This project demonstrates how to integrate the OpenWeather API into a Go application to retrieve weather information based on latitude and longitude coordinates.
Clone the repo:

## Navigate to the project directory:
cd weather-api

## Configuration
- `OPENWEATHER_API_KEY`: Your API key for accessing OpenWeather data.



You can set your api key environment variable either in your system environment or by creating a .env file in the project directory with the following content:
export OPENWEATHER_API_KEY=<your-api-key>


## Usage

Start the server:
go run .
The server will start running on http://localhost:8080.
Make a GET request to the /weather endpoint:
http://localhost:8080/weather?lat=<latitude>&lon=<longitude>
Replace <latitude> and <longitude> with the desired latitude and longitude coordinates.
Heres an example:
http://localhost:8080/weather?lat=37.7749&lon=-122.4194
The API will respond with a JSON object containing the weather information for the specified latitude and longitude.

## API Response
The API response is a JSON object with the following structure:
jsonCopy code{
  "weatherCondition": "Clear",
  "temperature": 25.5,
  "tempDescription": "hot"
}

weatherCondition: The main weather condition (e.g., Clear, Clouds, Rain).
temperature: The temperature in Celsius.
tempDescription: A description of the temperature (cold, moderate, or hot).

## Error Handling
The API handles the following error scenarios:

Missing latitude or longitude: If the LATITUDE or LONGITUDE environment variables are not set, the API will respond with a 400 Bad Request status code and an error message.
Missing API key: If the OPENWEATHER_API_KEY environment variable is not set, the API will respond with a 500 Internal Server Error status code and an error message.
Failed to fetch weather data: If there is an error while fetching data from the OpenWeather API, the API will respond with a 500 Internal Server Error status code and an error message.
Weather information not found: If the weather information is not available in the OpenWeather API response, the API will respond with a 404 Not Found status code and an error message.