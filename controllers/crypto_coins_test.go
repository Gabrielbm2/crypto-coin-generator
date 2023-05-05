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

func TestCreateCrypto(t *testing.T) {
	newCrypto := models.CreateCrypto{
		ID:        "",
		Name:      "Bitcoin",
		CreatedAt: time.Time{},
	}
	body, err := json.Marshal(newCrypto)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/cryptos", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateCrypto)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	assert.Equal(t, "", rr.Body.String())
}
