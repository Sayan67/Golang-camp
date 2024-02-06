package main

import "fmt"

func main() {
	var array [5]string
	array[0] = "Sayan Das"
	array[2] = "Sailen Das"
	array[3] = "Sadhana Das"
	array[1] = "Homieeee"
	fmt.Println("Family : ", array)
	fmt.Println("Length of family is : ", len(array))
}
