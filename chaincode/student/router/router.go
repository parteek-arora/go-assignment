package router

import (
	student "assignment/chaincode/student/chaincode"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"gopkg.in/go-playground/validator.v9"
)

type StudentModel = student.Student

var id int

// model for success response 200
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//model for error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

/**
@description add all the students router
@param  pointer to the router
*/
func AddStudentRoutes(router *mux.Router, test *testing.T, stub *shim.MockStub) {

	/**
	@description router for add new student
	@method post
	@path /students
	@param callback method with res and req
	@response with add student detail or error
	*/
	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		var studentData StudentModel
		err := json.NewDecoder(r.Body).Decode(&studentData)
		if err != nil {
			error := ErrorResponse{400, err.Error()}
			error.sendError(w)
			return
		}
		id = id + 1
		studentData.Id = id
		// //validation for the struct
		validate := validator.New()
		validationError := validate.Struct(studentData)
		if validationError != nil {
			id = id - 1
			error := ErrorResponse{400, validationError.Error()}
			error.sendError(w)
			return
		}
		//invoke chaincode method AddStudent
		newStudentJsonBytes, ccerror := Invoke(test, stub, "AddStudent", strconv.Itoa(studentData.Id), studentData.Name, studentData.Phone)
		if ccerror != nil {
			id = id - 1
			error := ErrorResponse{400, ccerror.(string)}
			error.sendError(w)
			return
		}
		newStudent := StudentModel{}
		//umarshal the data to a new ballot struct
		json.Unmarshal(newStudentJsonBytes, &newStudent)
		res := &SuccessResponse{200, "success", newStudent}
		res.sendStudentResponse(w)
	}).Methods("POST")

	/**
	@description router for edit student details
	@method put
	@path /students
	@param callback method with res and req
	@response with edited student detail or error
	*/
	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		var studentData StudentModel
		err := json.NewDecoder(r.Body).Decode(&studentData)
		if err != nil {
			error := ErrorResponse{400, err.Error()}
			error.sendError(w)
			return
		}
		//validate the request body
		validate := validator.New()
		validationError := validate.Struct(studentData)
		if validationError != nil {
			error := ErrorResponse{400, validationError.Error()}
			error.sendError(w)
			return
		}
		//invoke chaincode method  EditStudent to edit the students
		editStudentJsonBytes, ccerror := Invoke(test, stub, "EditStudent", strconv.Itoa(studentData.Id), studentData.Name, studentData.Phone)
		if ccerror != nil {
			id = id - 1
			error := ErrorResponse{400, ccerror.(string)}
			error.sendError(w)
			return
		}
		student := StudentModel{}
		//umarshal the data to a new ballot struct
		json.Unmarshal(editStudentJsonBytes, &student)
		res := &SuccessResponse{200, "success", student}
		res.sendStudentResponse(w)
	}).Methods("PUT")

	/**
	@description router for all the students list
	@method get
	@path /students
	@param callback method with res and req
	@response with complete student list or error
	*/
	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		//invoke ViewAllStudents method
		studentsBytes, ccerror := Invoke(test, stub, "ViewAllStudents")
		if ccerror != nil {
			error := ErrorResponse{400, ccerror.(string)}
			error.sendError(w)
			return
		}
		students := []StudentModel{}
		//umarshal the data to a new ballot struct
		json.Unmarshal(studentsBytes, &students)
		res := &SuccessResponse{200, "success", students}
		res.sendStudentResponse(w)
	}).Methods("GET")

	/**
	@description router to delete  student
	@method delete
	@path /students/id
	@param  callback method with res and req
	@response  deleted student detail or error
	*/
	router.HandleFunc("/students/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		//invoke chaincode method DeleteStudent
		_, ccerror := Invoke(test, stub, "DeleteStudent", params["id"])
		if ccerror != nil {
			error := ErrorResponse{400, ccerror.(string)}
			error.sendError(w)
			return
		}
		res := &SuccessResponse{200, "success", params["id"]}
		res.sendStudentResponse(w)
	}).Methods("DELETE")

	/**
	@description router get detail of particular student
	@method get
	@path /students/id
	@param  callback method with res and req
	@response  student detail or error
	*/
	router.HandleFunc("/students/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		//invoke method QueryStudent
		studentsBytes, ccerror := Invoke(test, stub, "QueryStudent", params["id"])
		if ccerror != nil {
			error := ErrorResponse{400, ccerror.(string)}
			error.sendError(w)
			return
		}
		student := StudentModel{}
		//umarshal the data to a new ballot struct
		json.Unmarshal(studentsBytes, &student)
		res := &SuccessResponse{200, "success", student}
		res.sendStudentResponse(w)
	}).Methods("GET")

}

/*common methos for invoking chaincode method
@desc change parameter to bytes
	  add fucntion name ass first paramaetr
@return bytes , interface for error
*/
func Invoke(test *testing.T, stub *shim.MockStub, function string, args ...string) ([]byte, interface{}) {

	cc_args := make([][]byte, 1+len(args))
	cc_args[0] = []byte(function)
	for i, arg := range args {
		cc_args[i+1] = []byte(arg)
	}
	result := stub.MockInvoke("000", cc_args)
	fmt.Println("result")
	fmt.Printf("%+v\n", result)
	fmt.Println("Call:    ", function, "(", strings.Join(args, ","), ")")
	fmt.Println("Code: ", result.Status)
	fmt.Println("Payload: ", string(result.Payload))
	if result.Status != shim.OK {
		// test.FailNow()
		return nil, result.Message
	} else {
		return result.Payload, nil
	}
}

// common method to send error response using pointer to fucntion
func (err *ErrorResponse) sendError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	payload, error := json.Marshal(err)
	if error != nil {
		log.Println(err)
	} else {
		w.Write(payload)
	}
}

//common method to send response
func (res *SuccessResponse) sendStudentResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	resData, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	} else {
		w.Write(resData)
	}
}
