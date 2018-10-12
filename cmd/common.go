package cmd

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/impactasaurus/soq-api/api"
	corsLib "github.com/rs/cors"
)

func MustSetup() Network {
	c := MustGetConfiguration()

	v1Handlers, err := api.NewV1()
	if err != nil {
		log.Fatal(err, nil)
	}

	cors := corsLib.New(corsLib.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:           86400,
	})

	r := mux.NewRouter()
	for _, h := range v1Handlers {
		r.Handle(h.Route, cors.Handler(h.Handler))
	}
	http.Handle("/", r)

	return c.Network
}
