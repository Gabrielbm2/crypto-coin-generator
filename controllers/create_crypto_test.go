package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"desafioKlever/models"
)

func TestCreateCryptos(t *testing.T) {
	// create a mock crypto to be sent in the request body
	newCrypto := models.CreateCrypto{
		ID:        "",
		Name:      "Bitcoin",
		CreatedAt: time.Time{},
	}
	body, err := json.Marshal(newCrypto)
	assert.NoError(t, err)

	// create a new request with the mock crypto as the request body
	req, err := http.NewRequest("POST", "/crypto-coins", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// create a mock response recorder
	rr := httptest.NewRecorder()

	// create a new handler and call it with the mock request and response
	handler := http.HandlerFunc(CreateCrypto)
	handler.ServeHTTP(rr, req)

	// assert that the response status code is 201 Created
	assert.Equal(t, http.StatusCreated, rr.Code)

	// assert that the response body is empty
	assert.Equal(t, "", rr.Body.String())
}
