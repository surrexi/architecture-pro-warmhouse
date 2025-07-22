package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	rand.New(rand.NewSource(rand.Int63()))

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		sensorID := r.URL.Query().Get("sensorId")

		// If no location is provided, use a default based on sensor ID
		if location == "" {
			switch sensorID {
			case "1":
				location = "Living Room"
			case "2":
				location = "Bedroom"
			case "3":
				location = "Kitchen"
			default:
				location = "Unknown"
			}
		}

		// If no sensor ID is provided, generate one based on location
		if sensorID == "" {
			switch location {
			case "Living Room":
				sensorID = "1"
			case "Bedroom":
				sensorID = "2"
			case "Kitchen":
				sensorID = "3"
			default:
				sensorID = "0"
			}
		}

		temp := rand.Float64()*10 + 18 // Температура от 18 до 28°C
		resp := TemperatureResponse{
			Value:       temp,
			Unit:        "°C",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "active",
			SensorID:    sensorID,
			SensorType:  "",
			Description: "Description",
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("encode failed: %v\n", err)
			return
		}
	})

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
