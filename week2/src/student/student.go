//place this file into GOPATH/src/customer

package student

import "fmt"

type Student struct {
	id    int
	name  string
	phone string
}

//method to add student
func AddStudent(studentsArray *[]Student, id int) {
	fmt.Println("********* Add New Student ***********")
	var name, phone string
	fmt.Println("Enter Student Name :-  ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter Student phone number :-  ")
	fmt.Scanf("%s", &phone)
	newStudent := Student{id, name, phone}
	*studentsArray = append(*studentsArray, newStudent)
	fmt.Println("********* New Student Added Successfully *********")
}

//method to edit the sudnt details
func EditStudent(studentsArray *[]Student) {
	fmt.Println("********* Edit Student's detail ***********")
	var id, index int
	var name, phone string
	fmt.Println("Enter id of student you want to edit :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range *studentsArray {
		if elem.id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		fmt.Println("## Invalid Id !! Please try again ##")
	} else {
		fmt.Println("Enter new Name :-  ")
		fmt.Scanf("%s", &name)
		fmt.Println("Enter new  phone number :-  ")
		fmt.Scanf("%s", &phone)
		(*studentsArray)[index-1].name = name
		(*studentsArray)[index-1].phone = phone
		fmt.Println("********* Details updated Successfully *********")
	}
}

//method to delete the student by id
func DeleteStudent(studentsArray *[]Student) {
	fmt.Println("********* Delete Student By Id ***********")
	var id, index int
	fmt.Println("Enter id of student you want to delete :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range *studentsArray {
		if elem.id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		fmt.Println("## Invalid Id !! Please try again ##")
	} else {
		index = index - 1
		//create new slice for maintaing the order
		*studentsArray = append((*studentsArray)[:index], (*studentsArray)[index+1:]...)
		fmt.Println("********* Student deleted Successfully *********")
	}
}

//view student detail by ID
func ViewStudentById(studentsArray *[]Student) {
	fmt.Println("********* View Student By Id ***********")
	var id, index int
	fmt.Println("Enter id of student you want to view :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range *studentsArray {
		if elem.id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		fmt.Println("## Invalid Id !! Please try again ##")
	} else {
		fmt.Println("Name of student is :- ", (*studentsArray)[index-1].name)
		fmt.Println("Phone Number of student is :- ", (*studentsArray)[index-1].phone)
	}
}

//view all students details
func ViewAllStudents(studentsArray *[]Student) {
	fmt.Println("********* List of all students ***********")
	fmt.Println("Total number of students :- ", len(*studentsArray))
	//view all data
	for _, elem := range *studentsArray {
		fmt.Println("| ", elem.id, " | ", elem.name, " | ", elem.phone, " |")
	}
	fmt.Println("********* end *********")
}
