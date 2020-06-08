package main

import (
	"fmt"
	"math"
	)

var number, count int = 2, 0

func main() {
	var i, n int
	fmt.Print("Enter the  Number : ")
	fmt.Scanln(&n)
	fmt.Println("The number entered is : ", n)
	fmt.Println("The prime numbers are :- ")
	
	// get the prime number
	for count <= n-1 {
		prime := true;
		for i = 2; i <= int(math.Sqrt(float64(number))); i++ {
			if number%i == 0 {
				prime = false;
				break
			}
		}
		if prime == true {
			count++
			fmt.Println(number)
		}
		number++
	}
}
