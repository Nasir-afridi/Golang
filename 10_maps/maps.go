package main

import (
	"fmt"
	"maps"
)

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

	// fixed length kk liye yane static map.
	a := map[string]int{"phones": 4, "price": 50}
	fmt.Println(a)

	// checking the element is in the map or not.

	_, hi := a["phones"] // underscore syyy ye zahir hota hai kkk ham phones ki value ko ignore krrhy hain .
	if hi {
		fmt.Println("hi all ok ")
	} else {
		fmt.Println("Not ok")
	}

	// checking that the maps are equal or not
	a1 := map[string]int{"phones": 4, "price": 50}
	a2 := map[string]int{"phones": 4, "price": 50}
	fmt.Println(maps.Equal(a1, a2))
}
