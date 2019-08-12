package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type weatherData struct {
	LocationName string `json:"location"`
	Weather      string `json:"weather"`
	Temperature  int    `json:"temp"`
	Celsius      bool   `json:"celsius,omitempty"`
	TempForecast []int  `json:"temp_forecast"`
}

func main() {
	weather := weatherData{
		LocationName: "Zzyzx",
		Weather:      "sunny",
		Temperature:  80,
		Celsius:      false,
		TempForecast: []int{76, 79, 85},
	}
	data, err := json.MarshalIndent(weather, "", "  ")
	if err != nil {
		log.Fatal("MarshalIndent failed:", err)
	}

	fmt.Println(string(data))

}
