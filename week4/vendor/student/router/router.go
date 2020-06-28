package router

import (
	"encoding/json"
	"log"
	"net/http"
	student "student/controller"

	"github.com/gorilla/mux"
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
func AddStudentRoutes(router *mux.Router) {

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
		//validation for the struct
		validate := validator.New()
		validationError := validate.Struct(studentData)
		if validationError != nil {
			id = id - 1
			error := ErrorResponse{400, validationError.Error()}
			error.sendError(w)
			return
		}
		newStudent := student.AddStudent(&studentData)
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
		editedStudent, errorMsg := student.EditStudent(&studentData)
		if errorMsg != nil {
			error := ErrorResponse{400, errorMsg.(string)}
			error.sendError(w)
			return
		}
		res := &SuccessResponse{200, "success", editedStudent.(StudentModel)}
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
		students := student.ViewAllStudents()
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
		deletedStudent, errorMsg := student.DeleteStudent(params["id"])
		if errorMsg != nil {
			error := ErrorResponse{400, errorMsg.(string)}
			error.sendError(w)
			return
		}
		res := &SuccessResponse{200, "success", deletedStudent.(StudentModel)}
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
		studentDetail, errorMsg := student.ViewStudentById(params["id"])
		if errorMsg != nil {
			error := ErrorResponse{400, errorMsg.(string)}
			error.sendError(w)
			return
		}
		res := &SuccessResponse{200, "success", studentDetail.(StudentModel)}
		res.sendStudentResponse(w)
	}).Methods("GET")

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
