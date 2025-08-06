package main

import "fmt"

// Slices : Dynamic arrays. ye bohot zyada use hotay hain go mein + ismein useful methods be hoty hain.

func main() {
	// uninatilized slice is nill.
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums)) // checking the length.

	// Make ka use krny syy hamesha memory allocate hoge chahay value ho ya na ho.
	// make ka use na krny sy memory tbb allocate hoge jbb ham kuch value daingy slice ko.

	// int kk bad length hai length kk bad capacity. capacity nahe daingy to jtne length hoge otne he capacity be hoge by default.
	var num = make([]int, 0, 5)
	fmt.Println(num)

	// capacity -> Maximum number of elements can fit.
	fmt.Println(cap(num))

	// append -> slices kk andr value add krny kk liye hota hai.
	num = append(num, 34, 54)
	fmt.Println((num))
	fmt.Println(cap(num))
	fmt.Println(len(num))

}
