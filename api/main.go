package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpPort int

func main() {
	flag.IntVar(&httpPort, "port", 8000, "HTTP Port")

	s := NewServer()
	log.Printf("Starting server on port %d\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), s.Router()))
}
