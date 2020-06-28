package controller

import (
	"strconv"
)

type Student struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=2,max=40"`
	Phone string `json:"phone" validate:"required,min=10,max=12"`
}

var StudentsArray []Student

/**
@description Add students in student array
@var int id
@param stdent model with name and phone
@return  Student (new added student structure)
*/
func AddStudent(student *Student) Student {
	StudentsArray = append(StudentsArray, *student)
	return *student
}

/**
@description edit the students in the student array
@param id , new name , new phone
@return interface for student model and error
*/
func EditStudent(student *Student) (interface{}, interface{}) {
	var index int
	for i, elem := range StudentsArray {
		if elem.Id == student.Id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		return nil, "invalid id"
	} else {
		StudentsArray[index-1].Name = student.Name
		StudentsArray[index-1].Phone = student.Phone
		return StudentsArray[index-1], nil
	}
}

/**
@description delete the student with id input
@param id string
@return interface for student model and error
*/

func DeleteStudent(paramId string) (interface{}, interface{}) {
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return nil, "invalid parmeter id"
	}
	var index int
	for i, elem := range StudentsArray {
		if elem.Id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		return nil, "invalid id"
	} else {
		index = index - 1
		deletedStudent := StudentsArray[index]
		StudentsArray = append((StudentsArray)[:index], (StudentsArray)[index+1:]...)
		return deletedStudent, nil
	}
}

/**
@description View the student with given id
@param id string
@return interface for student model and error
*/
func ViewStudentById(paramId string) (interface{}, interface{}) {
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return nil, "invalid parmeter id"
	}
	var index int
	for i, elem := range StudentsArray {
		if elem.Id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		return nil, "invalid id"
	} else {
		studentjson := StudentsArray[index-1]
		return studentjson, nil
	}
}

/**
@description view all students
@param --
@return Array of students
*/
func ViewAllStudents() []Student {
	return StudentsArray
}
