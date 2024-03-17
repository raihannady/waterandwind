package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go updateJSONFile()

	for {
		status := readJSONFile()
		waterStatus := getStatus(status.Water)
		windStatus := getStatus(status.Wind)

		fmt.Printf("Water: %d meter - Status: %s\n", status.Water, waterStatus)
		fmt.Printf("Wind: %d meter/detik - Status: %s\n", status.Wind, windStatus)

		time.Sleep(15 * time.Second)
	}
}

func updateJSONFile() {
	for {
		status := Status{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}

		file, err := os.Create("data.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		err = encoder.Encode(status)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// fmt.Println("Data updated successfully")

		time.Sleep(15 * time.Second)
	}
}

func readJSONFile() Status {
	var status Status

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return status
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&status)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return status
	}

	return status
}

func getStatus(value int) string {
	switch {
	case value < 5:
		return "Aman"
	case value >= 6 && value <= 8:
		return "Siaga"
	default:
		return "Bahaya"
	}
}
