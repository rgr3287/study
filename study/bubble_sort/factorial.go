package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// 팩토리얼 짜보기
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
