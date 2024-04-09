package main

import (
	"fill-asset-process-prototype/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/fill-users", routes.HandlerUserGet).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
