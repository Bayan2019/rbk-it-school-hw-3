package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.User.Create)
			r.Get("/", h.User.List)
			r.Get("/{id}", h.User.GetByID)
			r.Put("/{id}", h.User.Update)
			r.Delete("/{id}", h.User.Delete)
			r.Route("/{id}/cities", func(r chi.Router) {
				r.Post("/", h.City.Add2User)
				r.Get("/", h.City.ListOfUser)
				r.Delete("/{city_id}", h.City.DeleteFromUser)
			})
		})
	})

	return r
}
