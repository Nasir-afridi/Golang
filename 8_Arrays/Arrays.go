package main

import "fmt"

// Ham array declare krty waqt khud btaingy kk iski ktne length hai.
func main() {
	vals := [3]int{2, 3, 4} //creating array.
	fmt.Println(vals)

	var nums [4]int
	nums[0] = 1 // adding the value.

	fmt.Println((nums))
	fmt.Println(len(nums)) // get the array length
}
