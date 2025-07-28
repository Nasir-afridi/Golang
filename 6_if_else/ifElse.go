package main

import "fmt"

// go do not have ternary operator so we use if-else in it.
func main() {
	age := 16

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is kid")
	}

	//Operators
	var userRole = "Admin"
	var hasPermissions = true
	// Or
	if userRole == "Admin" || hasPermissions {
		fmt.Println("yes")
	} else if userRole == "Admin" && hasPermissions { // and operator.
		fmt.Println("yes")
	}

	// we can declare a variable inside the condition.
	if height := 5; height >= 6 {
		fmt.Println(" Perfect")
	} else if height >= 3 {
		fmt.Println("Not perfect")
	}
}
