// post.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func generateRandomData() Data {
	return Data{
		Wind:  rand.Intn(100) + 1,
		Water: rand.Intn(100) + 1,
	}
}

func sendDataToEndpoint(data Data) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
	return nil
}

func sendRandomData() {
	for {
		data := generateRandomData()
		fmt.Printf("Sending data: %+v\n", data)
		err := sendDataToEndpoint(data)
		if err != nil {
			fmt.Println("Error sending data:", err)
		}
		time.Sleep(15 * time.Second)
	}
}
