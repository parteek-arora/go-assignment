//place the file in the gopath src
package student

import ( 
	"fmt"
	"encoding/json"
)

type Student struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type StudentGroup struct {
	StudentsArray []Student
}


/**
 Add students in associated group reference to the called group struct
 @var int id
 @input name , phone
 @return  Student (new added student structure)
*/
func (group *StudentGroup) AddStudent(id int) Student{
	fmt.Println("********* Add New Student ***********")
	var name, phone string
	fmt.Println("Enter Student Name :-  ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter Student phone number :-  ")
	fmt.Scanf("%s", &phone)
	newStudent := &Student{id, name, phone};
	//show json data 
	studentjson, _ := json.MarshalIndent(newStudent , "", "")
    fmt.Println("JSON  --> " , string(studentjson))
	group.StudentsArray = append(group.StudentsArray, *newStudent)
	fmt.Println("********* New Student Added Successfully *********")
	return *newStudent;
}



/**
 Add students in associated group reference to the called group struct
 @var int id 
 @input name , phone data in the form of json string
 @return  int 0/1 -> error/success
*/
func (group *StudentGroup) AddStudentJSON(id int) int{
	fmt.Println("********* Add JSON for New Student ***********")
	var studentJSONString string
	fmt.Println("Enter Student JSON :-  ")
	fmt.Scanf("%s", &studentJSONString)
    fmt.Println("JSON  --> " , studentJSONString)
	studentInfo := &Student{};
	studentInfo.Id = id;
	err := json.Unmarshal([]byte(studentJSONString), studentInfo);
	if err != nil {
		fmt.Println(err)
		return 0;
	}else {
		group.StudentsArray = append(group.StudentsArray, *studentInfo);
		fmt.Println("********* New Student Added Successfully *********")
		return 1;
	}
	
}



/**
 edit the students details with reference to the group
 @var -- 
 @input id , new name , new phone
*/
func (group *StudentGroup) EditStudent() {
	fmt.Println("********* Edit Student's detail ***********")
	var id, index int
	var name, phone string
	fmt.Println("Enter id of student you want to edit :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range group.StudentsArray {
		if elem.Id == id {
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
		group.StudentsArray[index-1].Name = name
		group.StudentsArray[index-1].Phone = phone
		studentjson,_ := json.MarshalIndent((group.StudentsArray)[index-1],"","");
		fmt.Println("JSON Data of student is :- ", string(studentjson));
		fmt.Println("********* Details updated Successfully *********")
	}
}



/**
 delete the student with id input
 @var -- 
 @input id
*/

func (group *StudentGroup) DeleteStudent() {
	fmt.Println("********* Delete Student By Id ***********")
	var id, index int
	fmt.Println("Enter id of student you want to delete :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range group.StudentsArray {
		if elem.Id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		fmt.Println("## Invalid Id !! Please try again ##")
	} else {
		index = index - 1
		//create new slice for maintaing the order
		studentjson,_ := json.MarshalIndent((group.StudentsArray)[index],"","");
		fmt.Println("JSON Data of student is :- ", string(studentjson));
		group.StudentsArray = append((group.StudentsArray)[:index], (group.StudentsArray)[index+1:]...)
		fmt.Println("********* Student deleted Successfully *********")
	}
}


/**
 View the student with given id from the corresponding group
 @var -- 
 @input id
*/
func (group *StudentGroup) ViewStudentById() {
	fmt.Println("********* View Student By Id ***********")
	var id, index int
	fmt.Println("Enter id of student you want to view :-  ")
	fmt.Scanf("%v", &id)
	for i, elem := range group.StudentsArray {
		if elem.Id == id {
			index = i + 1
			break
		}
	}
	if index == 0 {
		fmt.Println("## Invalid Id !! Please try again ##")
	} else {
		fmt.Println("Name of student is :- ", (group.StudentsArray)[index-1].Name)
		fmt.Println("Phone Number of student is :- ", (group.StudentsArray)[index-1].Phone);
		//show inteded json data
		studentjson,_ := json.MarshalIndent((group.StudentsArray)[index-1],"","");
		fmt.Println("JSON Data of student is :- ", string(studentjson));
	}
}


/**
 get the students detail
 @var -- 
 @input --
 @return Array of students corresponding group
*/
func (group *StudentGroup) ViewAllStudents() [] Student{
	fmt.Println("********* List of all students ***********")
	fmt.Println("Total number of students :- ", len(group.StudentsArray))
	//view all data
	for _, elem := range group.StudentsArray {
		fmt.Println("| ", elem.Id, " | ", elem.Name, " | ", elem.Phone, " |")
	}
	//display array of json
	studentsArrayJSON,_:= json.Marshal(group.StudentsArray);
	fmt.Println("JSON Data of students is :- ", string(studentsArrayJSON));
	fmt.Println("********* end *********")
	return group.StudentsArray;
}
