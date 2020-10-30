package routes

import (
	"github.com/kirankumar26/smartfactory/controller"

	"github.com/gorilla/mux"
)

//	Define and Return application routes
func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)

	//routes.PathPrefix("/api/v1").Subrouter()

	routes.HandleFunc("/hello", controller.Hello).Methods("GET")
	routes.HandleFunc("/employee", controller.GetEmployees).Methods("GET")
	return routes
}
