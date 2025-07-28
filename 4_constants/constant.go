package main

import "fmt"

func main() {
	// ham isko change nahe krskty.
	const name = "Golang"
	fmt.Println(name)

	//we can group multiple constant together
	const (
		port = 5000
		host = "local host"
	)
	const (
		age      = 23
		language = "Golang"
	)
	fmt.Println(port, host)
	fmt.Println(age, language)
}
