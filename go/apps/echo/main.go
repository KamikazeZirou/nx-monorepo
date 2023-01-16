package main

import (
	"fmt"
	"github.com/kamikazezirou/nx-monorepo/go/libs/httplog"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}
	addr := net.JoinHostPort("", port)

	http.HandleFunc("/", httplog.Log(log.Default(), http.HandlerFunc(echoHandler)))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic("failed to listen")
	}

	srv := &http.Server{
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	body := r.Form.Get("body")
	_, _ = fmt.Fprint(w, body)
}
