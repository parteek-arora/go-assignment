package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	studentRouter "student/router"

	"github.com/gorilla/mux"
)

var httpPort int = 3000 //port to intialize the server

func main() {
	router := mux.NewRouter()

	//add all the routes of a student
	studentRouter.AddStudentRoutes(router)

	//customize the 404 handler to send json
	router.NotFoundHandler = http.HandlerFunc(notFound)
	log.Println("server started on port : ", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router)
	//error handling for server intialization (if running on same port)
	if err != nil {
		log.Fatal(err)
	}
}

/**
@description :- method to handle the 404 invalid route
@return   send error response to client
*/
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, "route not found")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	var errRes = &studentRouter.ErrorResponse{404, "Route not found"}
	payload, err := json.Marshal(errRes)
	if err != nil {
		log.Println(err)
	} else {
		w.Write(payload)
	}
}
