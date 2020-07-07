package main

import (
	"assignment/week2/student"
	"fmt"
)

var id int

type StudentDict = student.Student

var studentsArray []StudentDict

func main() {
	var optionSelected int
	fmt.Println("--------- Please select one option --------------------------")
	fmt.Println("Press 1 for add new student")
	fmt.Println("Press 2 for edit student's detail")
	fmt.Println("Press 3 for delete a student")
	fmt.Println("Press 4 for view student's detail by id")
	fmt.Println("Press 5 for view all the student")
	//dynmically select the option
	fmt.Scanln(&optionSelected)
	//switch case for the selected option
	switch optionSelected {
	case 1:
		//increment id and call the fucntion
		id++
		student.AddStudent(&studentsArray, id)
	case 2:
		student.EditStudent(&studentsArray) // edit student details
	case 3:
		student.DeleteStudent(&studentsArray) //delete student details
	case 4:
		student.ViewStudentById(&studentsArray) //view student details by id
	case 5:
		student.ViewAllStudents(&studentsArray) //view all sudents
	default:
		fmt.Println("Option selected is invalid. Please select again")
	}
	main()
}
