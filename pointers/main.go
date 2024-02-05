package main

import "fmt"

func main() {
	var ptr *int
	var val int
	val = 67
	fmt.Println("Value of empty pointer is : ", ptr)
	ptr = &val
	fmt.Println("Value at the location of the pointer is : ", *ptr)
	*ptr = *ptr + 5
	fmt.Println("Value after the addition is  : ", *ptr)
	fmt.Println("Address inside the pointer is : ", ptr)
}
