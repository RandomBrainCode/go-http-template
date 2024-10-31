// main_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal("Couldn't parse response: ", err)
	}

	expected := "Hello from Go!"
	if response.Message != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Message, expected)
	}
}
