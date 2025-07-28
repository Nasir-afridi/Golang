package main

import "fmt"

// for-> go kk andrr sirf yehe aik loop hai. ham while loop ko for loop sy he banaingy.
func main() {

	// while llop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // for loop
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue // iska mtlb hota hai kkk current value ko skip kro.
	// 	}
	// 	fmt.Println(i)
	// }

	// range concept
	for a := range 5 {
		fmt.Println(a) // 5 tkk yane 0,1,2,3,4 print kryga ye.
	}

}
