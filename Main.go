package main

import (
	"net/http"

	"github.com/kirankumar26/smartfactory/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Application has been started on : 8080")

	//creating instance of mux
	routes := routes.GetRoutes()
	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
