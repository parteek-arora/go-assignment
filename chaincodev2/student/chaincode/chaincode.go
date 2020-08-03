package chaincode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Student struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=2,max=40"`
	Phone string `json:"phone" validate:"required,min=10,max=12"`
}

//chaincode
type StudentChaincode struct {
	contractapi.Contract
}

/*
AddStudent method for tha chaincode
@desc chaincode to add a student
@return error
*/
func (stuchaincode *StudentChaincode) AddStudent(ctx contractapi.TransactionContextInterface, newStudent Student) error {

	//chaincode uniqe key
	chainCodeId := "STUDENT" + strconv.Itoa(newStudent.Id)

	// var student = Student{Id: id, Name: name, Phone: phone}
	//convert to json bytes
	studentJsonBytes, err := json.Marshal(newStudent)
	if err != nil {
		return fmt.Errorf("Failed to parse to json")
	}
	// add student using putState
	errPut := ctx.GetStub().PutState(chainCodeId, studentJsonBytes)
	if errPut != nil {
		return errPut
	}
	return nil
}

/*
QueryStudent method for tha chaincode
@desc chaincode to fetch a student details by give id
@return Student or error
*/
func (stuchaincode *StudentChaincode) QueryStudent(ctx contractapi.TransactionContextInterface, id string) (*Student, error) {
	//generate unique chaincode key
	chainCodeId := "STUDENT" + id
	//get the respective state data
	studentJsonBytes, errState := ctx.GetStub().GetState(chainCodeId)
	if errState != nil {
		return nil, errState
	}
	//chech if the data is present or not
	if len(studentJsonBytes) == 0 {
		return nil, fmt.Errorf("No Student Found")
	}
	//return bytes
	student := new(Student)
	//umarshal the data to a new ballot struct
	json.Unmarshal(studentJsonBytes, &student)
	return student, nil
}

/*
DeleteStudent method for tha chaincode
@desc chaincode to change the state status to deleted in the state table
@return  nil or error
*/

func (stuchaincode *StudentChaincode) DeleteStudent(ctx contractapi.TransactionContextInterface, id string) error {
	//generate unique chaincode key
	chainCodeId := "STUDENT" + id
	errState := ctx.GetStub().DelState(chainCodeId)
	if errState != nil {
		return errState
	}
	return nil
}

/*
EditStudent method for tha chaincode
@desc chaincode to change the state to edit the state data
@return  student or error
*/
func (stuchaincode *StudentChaincode) EditStudent(ctx contractapi.TransactionContextInterface, newStudent Student) error {

	//generate the unique id
	chainCodeId := "STUDENT" + strconv.Itoa(newStudent.Id)
	//get the data presnt on given state
	studentJsonBytes, errState := ctx.GetStub().GetState(chainCodeId)
	if errState != nil {
		return errState
	}
	if len(studentJsonBytes) == 0 {
		return fmt.Errorf("No Student Found")
	}

	//convert to jsonbytes
	newStudentJsonBytes, err := json.Marshal(newStudent)
	if err != nil {
		return fmt.Errorf("Failed to parse to json")
	}
	//change the state data u
	errPut := ctx.GetStub().PutState(chainCodeId, newStudentJsonBytes)
	if errPut != nil {
		return errPut
	}
	//return
	return nil
}

/*
ViewAllStudents method for tha chaincode
@desc chaincode to fetch all the state data between the given range
@return  students array or error
*/
func (stuchaincode *StudentChaincode) ViewAllStudents(ctx contractapi.TransactionContextInterface) (*[]Student, error) {
	//get the states data between the range
	keysIter, err := ctx.GetStub().GetStateByRange("STUDENT1", "student100")
	if err != nil {
		return nil, err
	}
	defer keysIter.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for keysIter.HasNext() {
		queryResponse, err := keysIter.Next()

		if err != nil {
			return nil, err
		}
		// // Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		// buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	students := []Student{}
	// 	//umarshal the data to a new student struct
	json.Unmarshal(buffer.Bytes(), &students)
	return &students, nil
}
