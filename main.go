package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Hello world h2c\n"
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, target)
	})
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(handler, h2s),
	}
	fmt.Print("Starting HTTP2 server\n")
	log.Fatal(server.ListenAndServe())
}
