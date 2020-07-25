package main

import (
	"log"
	"os"

	"github.com/NiceLabs/npm-registry-agent"
	"github.com/akrylysov/algnhsa"
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
	algnhsa.ListenAndServe(handler, nil)
}
