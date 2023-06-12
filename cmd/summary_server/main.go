package main

import (
	"net/http"
	"os"

	"github.com/eduaravila/stori-challenge/internal/errors/server"
	"github.com/eduaravila/stori-challenge/pkg/summary/ports"
	"github.com/eduaravila/stori-challenge/pkg/summary/service"
)

func main() {
	storageType := os.Getenv("STORAGE_TYPE")
	port := os.Getenv("SUMMARY_PORT")

	if port == "" {
		port = "8080"
	}

	app := service.NewApplication(storageType)

	server.RunHTTPServer("/api", ":"+port, func() http.Handler {
		return ports.Handler(ports.NewHTTPServer(app))
	})

}
