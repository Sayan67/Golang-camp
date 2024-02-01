package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the rating for the pizza : ")
	reader := bufio.NewReader(os.Stdin)
	inp, error := reader.ReadString('\n')
	newInput, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		fmt.Println(err)
	} else {

		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println("Actual ratng was : ", inp)
		}
		fmt.Println("New rating is : ", newInput+1)
	}
}
