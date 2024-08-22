package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"rest-workshop/src/router"
	"testing"
)

func TestCreationHandler(t *testing.T) {
	r := router.CreateRouter()
	data := struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Id    string  `json:"id"`
	}{Name: "something", Id: "miniId", Price: 0.323}
	marshalledData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/item", bytes.NewReader(marshalledData))

	req.Header.Add("content-type", "application/json")

	secondReq, _ := http.NewRequest("POST", "/item", bytes.NewReader(marshalledData))

	secondReq.Header.Add("content-type", "application/json")
	first := httptest.NewRecorder()
	second := httptest.NewRecorder()
	r.ServeHTTP(first, req)
	r.ServeHTTP(second, secondReq)

	result := struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Id    string  `json:"id"`
	}{}

	err := json.Unmarshal(first.Body.Bytes(), &result)

	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusCreated, first.Code)
	assert.Equal(t, http.StatusConflict, second.Code)
	assert.Equal(t, result, data)
}
