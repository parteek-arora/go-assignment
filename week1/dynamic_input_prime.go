package main

import "fmt"

var number, count int = 3, 2

func main() {
	var i, n int
	fmt.Print("Enter the  Number : ")
	fmt.Scanln(&n)
	fmt.Println("The number entered is : ", n)

	//logic to calculate the prime numbers
	if n >= 1 {
		fmt.Println("First ", n, " Prime Numbers Are")
		fmt.Println("2 ")
	}
	for count <= n {
		for i = 2; i < number; i++ {
			if number%i == 0 {
				break
			}
		}
		if i == number {
			count++
			fmt.Println(number)
		}
		number++
	}
}
