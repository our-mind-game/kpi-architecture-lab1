package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", getTime)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type TimeResponse struct {
	Time string `json:"time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	timeResponse := TimeResponse{
		Time: currentTime,
	}

	jsonResponse, err := json.Marshal(timeResponse)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
