package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"desafioKlever/models"
	"desafioKlever/testutils"
)

func TestReturnVotes(t *testing.T) {
	testutils.SetupTest()

	cryptoID := "BTC"
	votes := &models.Votes{
		Likes:    10,
		Dislikes: 5,
	}
	testutils.MockGetVotes(cryptoID, votes)

	req, err := http.NewRequest("GET", fmt.Sprintf("/votes/%s", cryptoID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ReturnVotes)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response ReturnVotesResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := ReturnVotesResponse{
		Likes:    votes.Likes,
		Dislikes: votes.Dislikes,
	}

	if !reflect.DeepEqual(&response, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response, expected)
	}

	testutils.ClearMockData()
}
