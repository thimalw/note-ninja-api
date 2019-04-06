package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/thimalw/note-ninja-api/user"
)

func routes(db *mgo.Database) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/user", user.Routes(db))
	})

	return r
}
