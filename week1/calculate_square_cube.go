package main

//import the package
import (
	"fmt"
	"squarecube"
)

func main() {
	var n int
	fmt.Print("Enter the  Number : ")
	fmt.Scanln(&n)
	fmt.Println("The number entered is : ", n);

	//use the method of thr imported package
	fmt.Println("The square of the number is : ", squarecube.Square(n))
	fmt.Println("The cube of the number is : ", squarecube.Cube(n))
}
