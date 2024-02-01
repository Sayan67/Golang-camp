package main

import (
	"fmt"
	"time"
)

func main() {
	//Present time calculation
	presentTime := time.Now()
	//This gives a unreadable format
	fmt.Println(presentTime)
	//Format specified by golang docs
	fmt.Println(presentTime.Format("15:04:05"))

	//Creating new date
	newDate := time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(newDate)
	//Formating the date according to golang's docs
	fmt.Println(newDate.Format("01-02-2006 Monday"))
}
