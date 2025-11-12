package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	}
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err

	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.OpenWeatherMapApiKey)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()
	var data weatherData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}
	return data, nil
}
func main() {
	// http.HandleFunc("/weather/", hello)
	r := mux.NewRouter()

	r.HandleFunc("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		city := params["city"]

		data, err := query(city)
		if err != nil {
			http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
