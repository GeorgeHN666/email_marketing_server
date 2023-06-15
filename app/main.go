package main

import (
	"log"

	"github.com/GeorgeHN/email-backend/app/http"
	"github.com/GeorgeHN/email-backend/app/router"
)

func main() {

	// HTTP SERVER
	config := http.Config{
		Addr:    3005,
		Timeout: 35,
		Handler: router.HandlerRoutes(),
	}

	log.Fatal(http.NewHTTPServer(config).Start())

}
