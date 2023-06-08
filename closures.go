package main

func AddN(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}
