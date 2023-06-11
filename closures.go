package main

import "fmt"

func AddN(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}

func RunClosuresDemo() {
	f := AddN(6)
	fmt.Println(f(51))
}
