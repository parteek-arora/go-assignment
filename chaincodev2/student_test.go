package main

import (
	studentChaincode "assignment/chaincodev2/student/chaincode"
	studentRouter "assignment/chaincodev2/student/router"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var httpPort int = 3000 //port to intialize the server

//main test fucntion to run
func TestMain(test *testing.T) {
	router := mux.NewRouter()
	//intialize the chaincode
	chaincode, err := contractapi.NewChaincode(new(studentChaincode.StudentChaincode))
	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		log.Fatal(err)
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
		log.Fatal(err)
	}
	//add all the routes of a student
	studentRouter.AddStudentRoutes(router, test, chaincode)

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
