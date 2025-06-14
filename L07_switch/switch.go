package main

import (
	"fmt"
	"time"
)


func main()  {
	// simple switch
	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println(("other"))
	}

	//multiple condition switch
    
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch s := i.(type) {
		case int:
			fmt.Println("its an intiger")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Print("its a boolean")
		default:
			fmt.Print("other", s)
		}
	}


	whoAmI(77.24)

	// whoAmI := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		fmt.Println("its an intiger")
	// 	case string:
	// 		fmt.Println("its a string")
	// 	case bool:
	// 		fmt.Print("its a boolean")
	// 	default:
	// 		fmt.Print("other")
	// 	}
	// }

}