package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping")
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", ping)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/", pong)

	go func() {
		log.Fatal(http.ListenAndServe(":29170", serverMuxA))
	}()

	log.Fatal(http.ListenAndServe(":29171", serverMuxB))
}
