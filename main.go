package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if fronEnv := os.Getenv("PORT"); fronEnv != "" {
		port = fronEnv
	}
	server := http.NewServeMux()

	server.HandleFunc("/", hello)
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))

}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "hello, world!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
