package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type
type Routes []Route

// NewRouter returns new router
func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		r := router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		if r.GetError() != nil {
			fmt.Printf("Error: %s (Route: %s)\n", r.GetError(), route.Pattern)
		}
	}
	return router
}
