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

	list := make([]WeatherForecast, 0, len(weatherForecasts))
	for _, value := range weatherForecasts {
		list = append(list, value)
	}

	json.NewEncoder(w).Encode(list)
}

func getByDateHandler(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	weatherForecast, found := weatherForecasts[date]

	if found {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(weatherForecast)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	temp := WeatherForecast{}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	weatherForecast := NewWeatherForecast(temp.Date, temp.TemperatureC, temp.Summary)
	weatherForecasts[temp.Date] = *weatherForecast

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

var weatherForecasts map[string]WeatherForecast = make(map[string]WeatherForecast)

func main() {
	setupMockData()

	router := mux.NewRouter()
	router.HandleFunc("/", getAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/{date}", getByDateHandler).Methods(http.MethodGet)
	router.HandleFunc("/", postHandler).Methods(http.MethodPost)

	fmt.Println("Http Listening at port localhost:8080")
	http.Handle("/", router)
	http.ListenAndServe(":8081", router)
}

func setupMockData() {
	summaries := []string{
		"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

	for i := 0; i < 6; i++ {
		d := time.Now().AddDate(0, 0, i).Format("2006-01-02")
		c := getRandomTemp()
		w := NewWeatherForecast(d, c, summaries[i])
		weatherForecasts[d] = *w
	}
}

func getRandomTemp() int {
	min := -20
	max := 55
	return rand.Intn(max-min) + min
}
