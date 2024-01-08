package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server: %s\n", err)
	}
}

func createServer(ink *Ink) (srv *http.Server) {
	r := gin.Default()
	setupRouter(ink, r)

	srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ink.env.Host, ink.env.Port),
		Handler: r,
	}

	go startServer(srv)

	return
}
