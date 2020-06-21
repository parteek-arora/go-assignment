package main

import (
	"fmt"
	"student"
)

var id int
//intialize structure with students slice and all the methods(like object of class)
var debutgroup = student.StudentGroup{make([] student.Student , 0)};


func main() {
	var optionSelected int
	fmt.Println("--------- Please select one option --------------------------")
	fmt.Println("Press 1 for add new student")
	fmt.Println("Press 2 for edit student's detail")
	fmt.Println("Press 3 for delete a student")
	fmt.Println("Press 4 for view student's detail by id")
	fmt.Println("Press 5 for view all the student")
	fmt.Println("Press 6 for add json data for student")
	//dynmically select the option
	fmt.Scanln(&optionSelected)
	//switch case for the selected option -> Access the method associate with the group
	switch optionSelected {
	case 1:
		//increment id and call the fucntion
		id++
		debutgroup.AddStudent(id)
	case 2:
		debutgroup.EditStudent() // edit student details
	case 3:
		debutgroup.DeleteStudent() //delete student details
	case 4:
		debutgroup.ViewStudentById() //view student details by id
	case 5:
		debutgroup.ViewAllStudents() //view all sudents
	case 6:
		id++
		debutgroup.AddStudentJSON(id) //add student json string
	default:
		fmt.Println("Option selected is invalid. Please select again")
	}
	main()
}
