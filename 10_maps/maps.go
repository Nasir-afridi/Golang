package main

import "fmt"

func main() {
	m := make(map[string]string) // initialized the map.

	// Adding the elements
	m["name"] = "golang"
	m["Area"] = "backend"
	m["phone"] = "0000"

	// get an element
	fmt.Println(m["name"], m["Area"], m["phone"])
	fmt.Println(len(m)) // checking the length

	delete(m, "phone") // delete the element from the map.
	fmt.Println(m)

	clear(m) // clear the map
	fmt.Println(m)
}
