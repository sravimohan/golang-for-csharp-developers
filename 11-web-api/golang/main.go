package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(weatherForecasts)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	temp := WeatherForecast{}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	weatherForecast := NewWeatherForecast(temp.Date, temp.TemperatureC, temp.Summary)
	weatherForecasts = append(weatherForecasts, *weatherForecast)

	response, err := json.Marshal(&weatherForecasts)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

var weatherForecasts = []WeatherForecast{}

func main() {
	setupMockData()

	router := mux.NewRouter()
	router.HandleFunc("/", getAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", postHandler).Methods(http.MethodPost)

	fmt.Println("Http Listening at port localhost:8080")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

func setupMockData() {
	summaries := []string{
		"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

	for i := 0; i < 6; i++ {
		d := dateOnly(time.Now().AddDate(0, 0, i))
		c := getRandomTemp()
		w := NewWeatherForecast(d, c, summaries[i])
		weatherForecasts = append(weatherForecasts, *w)
	}
}

func dateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func getRandomTemp() int {
	min := -20
	max := 55
	return rand.Intn(max-min) + min
}
