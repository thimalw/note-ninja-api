package user

import "github.com/go-chi/chi"

// Routes returns the REST routes for user
func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{userID}", User)
	return r
}
