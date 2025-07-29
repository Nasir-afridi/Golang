package main

import (
	"fmt"
	"time"
)

func main() {
	// switch
	i := 99
	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("five")

	case 99:
		fmt.Println("six")

	default:
		fmt.Println("other")
	}

	// Multiple condition switch
	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")

	default:
		fmt.Println("it's office day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) { // ye pata chalanay kk liye kkk ye kis type ka variable hai.

		case int:
			fmt.Println("its integer")

		case string:
			fmt.Println("it's string")

		case bool:
			fmt.Println("it's a boolean")

		default:
			fmt.Println("other")
		}
	}
	whoAmI(44)

	// fallthrough: condition check kiye bina next prr chala jata hai.
	num := 1

	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
		fallthrough
	case 4:
		fmt.Println("four")

	default:
		fmt.Println("not matched")
	}
}
