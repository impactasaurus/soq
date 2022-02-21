package cmd

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/impactasaurus/soq/api"
	"github.com/impactasaurus/soq/cache"
	"github.com/impactasaurus/soq/questionnaires"
	"github.com/impactasaurus/soq/search"
	corsLib "github.com/rs/cors"
)

func MustSetup() Network {
	cfg := MustGetConfiguration()

	qq, err := questionnaires.LoadAll()
	if err != nil {
		log.Fatal(err)
	}

	c, err := cache.New(qq)
	if err != nil {
		log.Fatal(err)
	}

	se, err := search.New(c)
	if err != nil {
		log.Fatal(err)
	}

	v1Handlers, err := api.NewV1(struct {
		*cache.Cache
		*search.Engine
	}{c, se})
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
