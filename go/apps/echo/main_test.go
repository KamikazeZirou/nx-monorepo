package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_echoHandler(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodGet, "/?body=foobar", nil)
	w := httptest.NewRecorder()

	// Act
	echoHandler(w, r)

	// Assert
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "foobar", string(body))
}
