package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLike(t *testing.T) {
	// create a new request
	req, err := http.NewRequest("POST", "/crypto/btc/like", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a mock response recorder
	rr := httptest.NewRecorder()

	// create a new handler and call it with the mock request and response
	handler := http.HandlerFunc(Like)
	handler.ServeHTTP(rr, req)

	// assert that the response status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDislike(t *testing.T) {
	// create a new request
	req, err := http.NewRequest("POST", "/crypto/btc/dislike", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a mock response recorder
	rr := httptest.NewRecorder()

	// create a new handler and call it with the mock request and response
	handler := http.HandlerFunc(Dislike)
	handler.ServeHTTP(rr, req)

	// assert that the response status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestGetVotes(t *testing.T) {
	// create a new request
	req, err := http.NewRequest("GET", "/crypto/btc/votes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a mock response recorder
	rr := httptest.NewRecorder()

	// create a new handler and call it with the mock request and response
	handler := http.HandlerFunc(GetVotes)
	handler.ServeHTTP(rr, req)

	// assert that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// assert that the response body is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
}
