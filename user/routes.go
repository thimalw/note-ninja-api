package user

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/thimalw/note-ninja-api/internal/utils"
)

// Routes returns the REST routes for user
func Routes(db *mgo.Database) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{userID}", utils.MakeHandler(User, db))
	return r
}
