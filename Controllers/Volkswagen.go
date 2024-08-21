package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var apiEndpoints = []string{
	"https://api1.example.com/data",
	"https://api2.example.com/data",
	"https://api3.example.com/data",
	"https://api4.example.com/data",
	"https://api5.example.com/data",
}

type ApiResponse struct {
	ErrorLog string `json:"error_log,omitempty"`
	WarnLog  string `json:"warn_log,omitempty"`
	Message  string `json:"message,omitempty"`
}

func FetchDataFromAPI(url string) (*ApiResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	var data ApiResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func AggregateResponses(apiEndpoints []string) map[string]int {
	logCount := map[string]int{"error": 0, "warn": 0, "info": 0}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range apiEndpoints {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			data, err := FetchDataFromAPI(url)
			if err != nil {
				log.Println("Failed to fetch data from", url, ":", err)
				return
			}

			mu.Lock()
			defer mu.Unlock()

			if data.ErrorLog != "" {
				logCount["error"]++
			} else if data.WarnLog != "" {
				logCount["warn"]++
			} else if data.Message != "" {
				logCount["info"]++
			}
		}(url)
	}

	wg.Wait()
	return logCount
}
