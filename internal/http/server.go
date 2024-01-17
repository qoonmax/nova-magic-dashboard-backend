package http

import (
	"net/http"
	"time"
)

func NewServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}
