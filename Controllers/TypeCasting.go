package Controllers

import (
	"encoding/json"
	"fmt"
)

func TypeCasting() {
	originalData := `
				{
					"boardInfo": {
						"data": "Random Data"
					},
					"things": [
						{
							"data": "Things Data 1"
						},
						{
							"data": "Things Data 2"
						},
                        {
							"data": "Things Data 3"
						},
                        {
							"data": "Things Data 4"
						}
					]
				}
				`
	newData := `
				{
                    "data": "New Things Data"
                }`

	/*
		Objective 1: Check the length of "things" in originalData
		Objective 2: Append the "newData" in "things" of originalData
		Objective 3: Verify by checking the length of "things"

	*/

	// start from here
	var orgData interface{}
	var new interface{}
	_ = json.Unmarshal([]byte(originalData), &orgData)
	_ = json.Unmarshal([]byte(newData), &new)
	things := orgData.(map[string]interface{})["things"].([]interface{})
	fmt.Println("Old", len(things))
	things = append(things, new)
	fmt.Println("New", len(things))
}
