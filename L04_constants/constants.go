package main

import "fmt"

const city = "delhi"

// not allowed
// country := "India"

func main() {
	const name = "golang"
	const age = 25
	

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

	fmt.Println(age)
	fmt.Println(city)
	fmt.Println(name)

	// print in same line
	fmt.Println(name, age, city)
}