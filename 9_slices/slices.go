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

	// empty slice created using literal notation
	nu := []int{}
	nu = append(nu, 10, 20, 30, 40, 50, 60)
	fmt.Println(nu)
	fmt.Println(len(nu))
	fmt.Println(cap(nu))

	// copy function.
	arr := make([]int, 0, 5)
	arr = append(arr, 4, 5, 6, 6, 6)
	fmt.Println(arr)

	arr2 := make([]int, len(arr))
	fmt.Println(arr2)

	copy(arr2, arr) // sabsy paihly osko rkhna hai jismein copy krna yane arr2 hia pher osay jissay copy krna hai arr.
	fmt.Println(arr, arr2)

	//Slices operator
	nmb := []int{1, 2, 4, 5, 6}
	fmt.Println(nmb[0:4]) // iska mtlb hai kk 0 index sy 4 index kk peechy tkk jtny elements hain sirf onko lo
	fmt.Println(nmb[:4])  // jbb left side prr kuch na ho to mtlb 0 sy start hoga 4 tkk.
	fmt.Println(nmb[0:])  // ismein ye hoga kk jo index ham daingy wahan sy end tkk sbb print kryga
}
