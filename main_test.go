package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"rest-workshop/src/database"
	"rest-workshop/src/router"
	"testing"
)

func TestCreationHandler(t *testing.T) {
	// Set up an in-memory database for testing
	dsn := "file::memory:?cache=shared"
	db, err := database.CreateDB(dsn)
	if err != nil {
		t.Fatalf("Failed to create in-memory database: %v", err)
	}

	// Create the router passing the test database
	r := router.CreateRouter(db)

	data := struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}{Name: "something", Price: 0.323}

	// Marshal the data to JSON format
	marshalledData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/item", bytes.NewReader(marshalledData))
	req.Header.Add("content-type", "application/json")

	// Create a response recorder for the request
	recorder := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(recorder, req)

	result := struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}{}

	// Decode the response
	err = json.Unmarshal(recorder.Body.Bytes(), &result)

	// Assertions
	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Verify that the 'Name' and 'Price' fields match
	assert.Equal(t, result.Name, data.Name)
	assert.Equal(t, result.Price, data.Price)
}
