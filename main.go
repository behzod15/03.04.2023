package main

import "fmt"

func main () {
	var num1, num2 int
	fmt.Println("Please enter first number of the range:")
	fmt.Scan(&num1)
	fmt.Println("Please enter second number of the range:")
	fmt.Scan(&num2)
	for i := num1; i <= num2; i++ {
		for j := 1; j <= i; j++{
			fmt.Println(i);
		}
	}
}