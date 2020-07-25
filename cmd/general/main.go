package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NiceLabs/npm-registry-agent"
)

func main() {
	var (
		token = os.Getenv("GITHUB_TOKEN")
		host  = os.Getenv("GITHUB_HOST")
	)
	if token == "" {
		log.Fatal("GITHUB_TOKEN is unset")
	} else if host == "" {
		log.Fatal("GITHUB_HOST is unset")
	}
	handler := agent.NewReverseProxy(host, token)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
