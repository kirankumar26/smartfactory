package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kirankumar26/smartfactory/service"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello I am Employee controller")
	result := service.GetEmployees()
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, result)
	json.NewEncoder(w).Encode(result)
}
