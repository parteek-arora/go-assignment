package chaincode

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Student struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=2,max=40"`
	Phone string `json:"phone" validate:"required,min=10,max=12"`
}

//chaincode
type StudentChaincode struct {
}

//Init implemented by StudentChaincode
func (t *StudentChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//Invoke implemented by StudentChaincode
func (stuchaincode *StudentChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//fc is the called fucntions
	fc, args := stub.GetFunctionAndParameters()
	if fc == "AddStudent" {
		return stuchaincode.AddStudent(stub, args)
	} else if fc == "ViewAllStudents" {
		return stuchaincode.ViewAllStudents(stub, args)
	} else if fc == "QueryStudent" {
		return stuchaincode.QueryStudent(stub, args)
	} else if fc == "DeleteStudent" {
		return stuchaincode.DeleteStudent(stub, args)
	} else if fc == "EditStudent" {
		return stuchaincode.EditStudent(stub, args)
	}
	return shim.Error("Called function is not defined in the chaincode ")
}

/*
AddStudent method for tha chaincode
@desc chaincode to add a student
@return bytes Student or error
*/
func (stuchaincode *StudentChaincode) AddStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//check the length of the arguments
	if len(args) != 3 {
		return shim.Error("3 argument required [id , name , phonenumber]")
	}
	//check if the id is valid or not
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Invalid student id")
	}
	//chaincode uniqe key
	chainCodeId := "STUDENT" + args[0]

	var student = Student{Id: id, Name: args[1], Phone: args[2]}
	//convert to json bytes
	studentJsonBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Failed to parse to json")
	}
	// add student using putState
	errPut := stub.PutState(chainCodeId, studentJsonBytes)
	if errPut != nil {
		return shim.Error(errPut.Error())
	}
	return shim.Success(studentJsonBytes)
}

/*
QueryStudent method for tha chaincode
@desc chaincode to fetch a student details by give id
@return bytes of Student or error
*/
func (stuchaincode *StudentChaincode) QueryStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//id is required
	if len(args) != 1 {
		return shim.Error("1 argument required [id]")
	}
	//generate unique chaincode key
	chainCodeId := "STUDENT" + args[0]
	//get the respective state data
	studentJsonBytes, errState := stub.GetState(chainCodeId)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	//chech if the data is present or not
	if len(studentJsonBytes) == 0 {
		return shim.Error("No Student Found")
	}
	//return bytes
	return shim.Success(studentJsonBytes)
}

/*
DeleteStudent method for tha chaincode
@desc chaincode to change the state status to deleted in the state table
@return  nil or error
*/
func (stuchaincode *StudentChaincode) DeleteStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//id to delete the student is required
	if len(args) != 1 {
		return shim.Error("1 argument required [id]")
	}
	//generate unique chaincode key
	chainCodeId := "STUDENT" + args[0]
	errState := stub.DelState(chainCodeId)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	//return
	return shim.Success(nil)
}

/*
EditStudent method for tha chaincode
@desc chaincode to change the state to edit the state data
@return  byte of student or error
*/
func (stuchaincode *StudentChaincode) EditStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//check if 3 args are present
	if len(args) != 3 {
		return shim.Error("3 argument required [id , name , phonenumber]")
	}
	// check or the id
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Invalid student id")
	}
	//generate the unique id
	chainCodeId := "STUDENT" + args[0]
	//get the data presnt on given state
	studentJsonBytes, errState := stub.GetState(chainCodeId)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	if len(studentJsonBytes) == 0 {
		return shim.Error("No Student Found")
	}

	var student = Student{Id: id, Name: args[1], Phone: args[2]}
	//convert to jsonbytes
	newStudentJsonBytes, err := json.Marshal(student)
	if err != nil {
		return shim.Error("Failed to parse to json")
	}
	//change the state data u
	errPut := stub.PutState(chainCodeId, newStudentJsonBytes)
	if errPut != nil {
		return shim.Error(errPut.Error())
	}
	//return bytes of new student data
	return shim.Success(newStudentJsonBytes)
}

/*
ViewAllStudents method for tha chaincode
@desc chaincode to fetch all the state data between the given range
@return  byte of students array or error
*/
func (stuchaincode *StudentChaincode) ViewAllStudents(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//get the states data between the range
	keysIter, err := stub.GetStateByRange("STUDENT1", "student100")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer keysIter.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for keysIter.HasNext() {
		queryResponse, err := keysIter.Next()

		if err != nil {
			return shim.Error(err.Error())
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
	// fmt.Printf("- queryAllStudents:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}
