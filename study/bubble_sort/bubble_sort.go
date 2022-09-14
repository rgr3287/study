package main

import "fmt"

func main() {
	a := [5]int{15, 1, 34, 6, 61}
	fmt.Println(&a)
	fmt.Println("오름 차순")
	fmt.Println(ascBubbleSort(a))
	fmt.Println("내림 차순")
	fmt.Println(descBubbleSort(a))
	fmt.Println("선택 정렬")
	fmt.Println(seletedSort(a))
	fmt.Println("선택정렬 최적화")
}

// 오름 차순 빈칸 채우기
func ascBubbleSort(a [5]int) [5]int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 내림 차순 변경
func descBubbleSort(a [5]int) [5]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 선택정렬
func seletedSort(a [5]int) [5]int {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}
