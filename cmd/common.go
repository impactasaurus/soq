package cmd

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/impactasaurus/soq-api/api"
	"github.com/impactasaurus/soq-api/cache"
	"github.com/impactasaurus/soq-api/questionnaires"
	corsLib "github.com/rs/cors"
)

func MustSetup() Network {
	cfg := MustGetConfiguration()

	qq, err := questionnaires.LoadDirectory(cfg.Path.Questionnaires)
	if err != nil {
		log.Fatal(err)
	}

	c, err := cache.New(qq)
	if err != nil {
		log.Fatal(err)
	}

	v1Handlers, err := api.NewV1(c)
	if err != nil {
		log.Fatal(err)
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

	return cfg.Network
}
