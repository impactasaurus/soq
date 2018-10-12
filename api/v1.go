package api

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/impactasaurus/soq-api/status"
)

type v1 struct {
	status *status.Provider
}

type RouteHandler struct {
	Route   string
	Handler http.Handler
}

// NewV1 returns a set of RouteHandler which serve V1 of the API
func NewV1() ([]RouteHandler, error) {
	v := &v1{
		status: status.New(),
	}
	return []RouteHandler{{
		Route:   "/v1/status",
		Handler: v.statusHandler(),
	}, {
		Route:   "/v1/playground",
		Handler: handler.Playground("GraphQL playground", "/v1/query"),
	}, {
		Route:   "/v1/query",
		Handler: handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})),
	}}, nil
}
