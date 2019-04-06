package utils

import (
	"net/http"

	"github.com/globalsign/mgo"
)

// MakeHandler encloses a REST route handler as a http.HandlerFunc
func MakeHandler(fn func(http.ResponseWriter, *http.Request, *mgo.Database), db *mgo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
