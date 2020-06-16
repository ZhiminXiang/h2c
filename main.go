package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http2.Server{}
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal("Fail to listen to " + port)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Fail to accept connection")
		}
		server.ServeConn(conn, &http2.ServeConnOpts{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello word: path %v\n", r.URL.Path)
			}),
		})
	}
}
