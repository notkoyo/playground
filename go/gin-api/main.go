package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MMR struct {
	Rank string `json:"currenttierpatched"`
	RR   int    `json:"ranking_in_tier"`
}

func main() {
	router := gin.Default()

	router.GET("/rank/:region/:name/:tag", func(c *gin.Context) {
		rankCall(c.region, c.name, c.tag)
	})

	router.Run("localhost:8080")
}

func rankCall(region string, name string, tag string) {
	url := `https://api.henrikdev.xyz/valorant/v1/mmr/%v/%v/%v?api_key=HDEV-5cdfb84f-9133-4ce3-a27b-db30845e6f17`

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response status:", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Extract the "data" part
	data, ok := result["data"]
	if !ok {
		fmt.Println("Error: 'data' key not found in the JSON response")
		return
	}

	// Convert the "data" part to JSON bytes
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling 'data' to JSON:", err)
		return
	}

	// Define a variable of type Person
	var mmr MMR

	// Unmarshal the JSON data into the person variable
	err = json.Unmarshal(dataBytes, &mmr)
	if err != nil {
		fmt.Println("Error unmarshaling 'data' JSON:", err)
		return
	}

	fmt.Printf("Rank: %s\nRR: %v\n", mmr.Rank, mmr.RR)
}
