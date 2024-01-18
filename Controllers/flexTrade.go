package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CityInfo struct {
	Name    string `json:"name"`
	Weather string `json:"weather"`
	Status  struct {
		Wind     string `json:"Wind"`
		Humidity string `json:"Humidity"`
	} `json:"status"`
}

func WeatherStation(keyword string) ([]string, error) {
	fmt.Println("hii")
	url := "https://jsonmock.hackerrank.com/api/weather/search?name="
	req, err := http.NewRequest("GET", url+keyword, nil)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Extract and format the information
	var result []string
	//var cities []CityInfo
	if data, ok := response["data"].([]interface{}); ok {

		for _, cityData := range data {
			if city, ok := cityData.(map[string]interface{}); ok {
				name := city["name"].(string)
				weather := city["weather"].(string)
				status := city["status"].([]interface{})
				wind := status[0].(string)
				humdity := status[1].(string)
				res := fmt.Sprintf("%s,%s,%s,%s", name, wind[6:7], humdity[9:12], weather[:2])
				result = append(result, res)
			}

		}
	}

	return result, nil
}
