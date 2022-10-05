package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(fibonacci(7))
}

// 팩토리얼 짜보기
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// 피보나치 수열
func fibonacci(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 하노이의 탑
func hanoiTop() {
	//a := [3]int{1, 2, 3}
	//var b []int
	//var c []int
	//	if a[0] > a[1] {
	//	}
	//
	//}
}
