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

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Like)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDislike(t *testing.T) {
	req, err := http.NewRequest("POST", "/crypto/btc/dislike", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Dislike)
	handler.ServeHTTP(rr, req)

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

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetVotes)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
}
