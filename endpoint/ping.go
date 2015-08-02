package endpoint

import (
	"fmt"
	"github.com/fake_validator/api"
	"net/http"
)

// Ping endpoint
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PONG")
}

// HealthCheck endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	result, err := api.HealthCheck()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, result)
}
