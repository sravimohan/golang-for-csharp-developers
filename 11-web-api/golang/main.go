package main

import (
	"fmt"
	"net/http"
	"youtube/golang-for-csharp-developers/11-web-api/golang/weather"

	"github.com/gorilla/mux"
)

func main() {
	weather.SetupMockData()

	router := mux.NewRouter()
	router.HandleFunc("/", weather.GetAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/{date}", weather.GetByDateHandler).Methods(http.MethodGet)
	router.HandleFunc("/", weather.PostHandler).Methods(http.MethodPost)

	fmt.Println("Http Listening at port localhost:8080")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
