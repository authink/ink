package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router"
)

func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server: %s\n", err)
	}
}

func createServer(ink *core.Ink) (srv *http.Server) {

	srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ink.Env.Host, ink.Env.Port),
		Handler: router.SetupRouter(ink),
	}

	go startServer(srv)

	return
}
