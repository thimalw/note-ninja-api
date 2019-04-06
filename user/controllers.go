package user

import (
	"net/http"

	"github.com/globalsign/mgo"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// User renders the details of a user specified by userID URL parameter in JSON
func User(w http.ResponseWriter, r *http.Request, db *mgo.Database) {
	userID := chi.URLParam(r, "userID")
	helloWorld := "Hello user " + userID
	render.JSON(w, r, helloWorld)
}
