// main.go
package main

import (
	"fmt"
	"time"
)

func getStatus(value int, safeThreshold int, alertThreshold int) string {
	if value <= safeThreshold {
		return "safe"
	} else if value <= alertThreshold {
		return "alert"
	}
	return "danger"
}

func main() {
	go sendRandomData()

	for {
		data := generateRandomData()

		waterStatus := getStatus(data.Water, 5, 8)
		windStatus := getStatus(data.Wind, 6, 15)

		fmt.Printf("water_value: %d\n", data.Water)
		fmt.Printf("wind_value: %d\n", data.Wind)
		fmt.Printf("water_status: %s\n", waterStatus)
		fmt.Printf("wind_status: %s\n\n", windStatus)

		time.Sleep(15 * time.Second)
	}
}
