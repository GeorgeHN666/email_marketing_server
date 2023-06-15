package http

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPServer struct {
	config Config
}

func NewHTTPServer(config Config) *HTTPServer {
	return &HTTPServer{
		config: config,
	}
}

func (s *HTTPServer) Start() error {

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", s.config.Addr),
		ReadTimeout:       time.Duration(s.config.Timeout) + 60*time.Second,
		WriteTimeout:      time.Duration(s.config.Timeout) + 60*time.Second,
		IdleTimeout:       time.Duration(s.config.Timeout) * time.Second,
		ReadHeaderTimeout: time.Duration(s.config.Timeout) + 25*time.Second,
		Handler:           s.config.Handler,
	}

	fmt.Printf("HTTP server starting at port: %d", s.config.Addr)

	return srv.ListenAndServe()
}
