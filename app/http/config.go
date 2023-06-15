package http

import "net/http"

type Config struct {
	Addr    int
	Handler http.Handler
	Timeout int
}
