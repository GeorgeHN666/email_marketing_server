package router

import (
	"net/http"

	"github.com/GeorgeHN/email-backend/app/handlers"
	"github.com/GeorgeHN/email-backend/app/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func HandlerRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"ZCS-23", "Accept", "Content-Type"},
		MaxAge:         350,
	}))

	mux.Post("/admin", handlers.NewAdmin)
	mux.Post("/client", handlers.NewClientEP)
	mux.Post("/campaing", handlers.InsertCampaingEP)
	mux.Post("/schedule", handlers.InsertScheduleEP)

	mux.Get("/img", middlewares.Ema(handlers.ServeImage))

	return mux
}
