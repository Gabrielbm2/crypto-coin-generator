package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Gabrielbm2/desafioKlever/models"
	"github.com/Gabrielbm2/desafioKlever/services/mocks"
)

func TestGetVotesByCryptoID(t *testing.T) {
	mockService := new(mocks.VoteService)

	controller := NewVotesController(mockService)

	vote := &models.Vote{
		ID:        1,
		CryptoID:  "BTC",
		Direction: "UP",
	}
	votes := []*models.Vote{vote}

	cryptoID := "BTC"

	mockService.On("GetVotesByCryptoID", cryptoID).Return(votes, nil)

	req, err := http.NewRequest("GET", "/crypto-coins/"+cryptoID+"/votes", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetVotesByCryptoID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedResponse := `[{"ID":1,"CryptoID":"BTC","Direction":"UP"}]`
	assert.Equal(t, expectedResponse, rr.Body.String())

	mockService.AssertCalled(t, "GetVotesByCryptoID", cryptoID)
}

func NewVotesController(mockService invalid type) {
	panic("unimplemented")
}
