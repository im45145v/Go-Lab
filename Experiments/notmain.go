//post request with bearer token
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace with your API endpoint URL
	apiUrl := "https://external-api.com/data"

	// Replace with your bearer token
	bearerToken := "your_bearer_token"

	// Data to send in the POST request body (optional)
	data := map[string]string{
		"field1": "value1",
		"field2": "value2",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set content type header (optional, might be inferred)
	req.Header.Set("Content-Type", "application/json")

	// Add authorization header with bearer token
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Process the response body (e.g., parse JSON)
	var responseData map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	fmt.Println("Response data:", responseData)

	// Use the retrieved data for further processing
	// ... your processing logic based on responseData ...
}
