package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i <= 10; i = i + 2 {
		sum = sum + i
	}
	fmt.Println(sum)
}
