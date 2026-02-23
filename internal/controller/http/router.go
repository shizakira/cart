package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	carthandlers "github.com/shizakira/cart/internal/controller/http/handlers"
	"github.com/shizakira/cart/pkg/zerochi"

	"github.com/shizakira/cart/internal/usecase"
)

func Router(uc *usecase.Cart) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(zerochi.Logger(log.Logger))
	r.Use(middleware.Recoverer)

	handlers := carthandlers.NewHandlers(uc)

	r.Route("/user/{user_id}/cart", func(r chi.Router) {
		r.Post("/{sku_id}", handlers.AddItem)
		r.Delete("/{sku_id}", handlers.RemoveItem)
		r.Delete("/", handlers.ClearCart)
		r.Get("/", handlers.GetItems)
	})

	return r
}
