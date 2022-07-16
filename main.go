package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AndreSC1218/GO_REST_API/routes"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
