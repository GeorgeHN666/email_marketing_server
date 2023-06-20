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
	mux.Post("/new_client", handlers.NewClientEP)
	mux.Post("/new_campaing", handlers.InsertCampaingEP)
	mux.Post("/new_schedule", handlers.InsertScheduleEP)
	mux.Get("/del_campaing", handlers.DeleteCampaingEP)

	mux.Post("/store_template", handlers.StoreTemplate)
	mux.Post("/store_audience", handlers.StoreAudience)

	mux.Get("/clients", handlers.GetClients)
	mux.Get("/client", handlers.GetClient)

	mux.Get("/campaings", handlers.GetCampaings)

	mux.Get("/schedules", handlers.GetSchedules)
	mux.Get("/del_schedules", handlers.DeleteScheduleEP)
	mux.Get("/del_client", handlers.DeleteClientEP)

	mux.Get("/img", middlewares.Ema(handlers.ServeImage))

	return mux
}
