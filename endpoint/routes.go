package endpoint

import (
	. "github.com/fake_validator/router"
)

// EndpointRoutes returns array of routes
var EndpointRoutes = Routes{
	Route{
		Name:        "Ping",
		Method:      "GET",
		Pattern:     "/ping",
		HandlerFunc: Ping,
	},
	Route{
		Name:        "Healthcheck",
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: HealthCheck,
	},
	Route{
		Name:        "NumberValidator",
		Method:      "GET",
		Pattern:     "/validate/number/{number}",
		HandlerFunc: ValidateNumber,
	},
}
