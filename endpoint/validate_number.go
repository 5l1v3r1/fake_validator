package endpoint

import (
	"fmt"
	"github.com/fake_validator/api"
	"github.com/gorilla/mux"
	"net/http"
)

// ValidateNumber endpoint
func ValidateNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	validator := api.NewNumberValidator(number)
	isValid, err := validator.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	fmt.Fprint(w, isValid)
}
