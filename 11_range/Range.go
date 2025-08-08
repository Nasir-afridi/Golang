package main

import (
	"fmt"
)

// Range -> iteration over data structure.
func main() {
	sum := 0
	nmb := []int{1, 2, 4}
	for _, num := range nmb { // _, iska mtlb hai kk hamain index nahe chaye ham osko ignore krrhy hain.
		sum = sum + num
	}
	fmt.Println(sum)

	// iteration over map.
	m := map[string]string{"fname": "nasir", "lname": "khan"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//iteration on the string to see the index and unicode code of the string.
	for a, b := range "Nasir" {
		fmt.Println(a, b)         // printing the unicode of each letter
		fmt.Println(a, string(b)) // printing the index with string Nasir
	}

}
