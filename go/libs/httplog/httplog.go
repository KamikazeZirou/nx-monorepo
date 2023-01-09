package httplog

import (
	"log"
	"net/http"
)

func Log(l *log.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Printf("%s: %s\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	}
}
