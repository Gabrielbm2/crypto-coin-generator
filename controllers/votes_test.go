package controllers_test

import (
	"desafioKlever/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVoteController(t *testing.T) {
	t.Run("Test Like", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/posts", nil)
		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := controllers.Like
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
		}

		var posts []Post

		if err := json.NewDecoder(rr.Body).Decode(&posts); err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}

		resultTotal := len(posts)
		expectedTotal := 100

		if resultTotal != expectedTotal {
			t.Errorf("Expected: %d. Got: %d.", expectedTotal, resultTotal)
		}
	})
}
