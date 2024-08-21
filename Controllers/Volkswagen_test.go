package Controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAggregateResponses(t *testing.T) {
	// Mock API responses
	mockAPIs := []struct {
		statusCode int
		response   ApiResponse
	}{
		{http.StatusOK, ApiResponse{ErrorLog: "error details", Message: "some error msg"}},
		{http.StatusOK, ApiResponse{ErrorLog: "error details", Message: "some error msg"}},
		{http.StatusOK, ApiResponse{WarnLog: "warn details", Message: "some error msg"}},
		{http.StatusOK, ApiResponse{Message: "info msg"}},
		{http.StatusOK, ApiResponse{Message: "info msg"}},
	}

	// Create a test server for each mock API
	var apiURLs []string
	for _, mock := range mockAPIs {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(mock.statusCode)
			json.NewEncoder(w).Encode(mock.response)
		}))
		defer ts.Close()
		apiURLs = append(apiURLs, ts.URL)
	}

	// Test the aggregation function
	oldEndpoints := apiEndpoints
	apiEndpoints = apiURLs
	defer func() { apiEndpoints = oldEndpoints }()

	result := AggregateResponses(apiEndpoints)

	expected := map[string]int{"error": 2, "warn": 1, "info": 2}
	if result["error"] != expected["error"] || result["warn"] != expected["warn"] || result["info"] != expected["info"] {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
