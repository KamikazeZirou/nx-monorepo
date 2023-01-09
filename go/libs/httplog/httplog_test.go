package httplog

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLog(t *testing.T) {
	// Arrange
	buf := bytes.NewBuffer([]byte{})
	l := log.New(buf, "", log.LstdFlags)
	var next http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	f := Log(l, next)

	// Act
	r := httptest.NewRequest(http.MethodGet, "/?body=foobar", nil)
	w := httptest.NewRecorder()
	f.ServeHTTP(w, r)

	// Assert
	assert.Contains(t, buf.String(), "GET: /?body=foobar")
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
