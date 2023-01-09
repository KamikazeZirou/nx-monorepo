package main

import (
	"fmt"
	"github.com/kamikazezirou/nx-monorepo/go/libs/httplog"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	addr := net.JoinHostPort("", port)

	http.HandleFunc("/", httplog.Log(log.Default(), http.HandlerFunc(echoHandler)))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	body := r.Form.Get("body")
	_, _ = fmt.Fprint(w, body)
}
