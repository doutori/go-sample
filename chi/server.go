package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Route("/members", func(r chi.Router) {
		r.Get("/", MemberIndex)
	})

	log.Println("Starting app")

	http.ListenAndServe(":8081", r)
}

// Member of type
type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MemberIndex ...
func MemberIndex(w http.ResponseWriter, r *http.Request) {
	members := []Member{
		{ID: 1, Name: "name1"},
		{ID: 2, Name: "name2"},
	}

	if err := json.NewEncoder(w).Encode(members); err != nil {
		panic(err)
	}
}
