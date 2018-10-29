package main

import (
	"fmt"
	"github.com/fake_validator/endpoint"
	"github.com/fake_validator/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter(endpoint.EndpointRoutes)

	fmt.Printf("Running on :8080 ...\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
