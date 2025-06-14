package main

import "fmt"

func main() {
	// var name string = "golang"

	// infer type
    // var name = "golang"

	var age = 25

	// shorthand syntax
	name := "golang"

    // type is necessary as variable is only defined and value is not assigned
	var city string
    city = "Delhi"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(city)
}