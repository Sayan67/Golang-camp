package main

import "fmt"

// Global variable
var globalVariable = "Hey global"

func main() {

	//Explicit type declaration
	var myFirstName string = "Sayan"
	fmt.Print(myFirstName)

	//implicit type declaration
	var mySecindName = "Das"
	fmt.Println(mySecindName)

	//Calling global variable
	fmt.Println(globalVariable)

	fmt.Printf("Global varible type is  : %T\n", globalVariable)

	//another type of variable declaration
	//this type of declaration is only possible inside a function
	specialVatr := 30000
	fmt.Println(specialVatr)

}
