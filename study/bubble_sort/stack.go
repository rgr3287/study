package main

import "fmt"

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("%d poped from stack\n", s.Pop())
}

type Stack []interface{}

//IsEmpty - 스택이 비어있는지 확인하는 함수
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Push - 스택에 값을 추가하는 함수.
func (s *Stack) Push(data interface{}) {
	*s = append(*s, data) // 스택 끝(top)에 값을 추가함.
	fmt.Printf("%d pushed to stack\n", data)
}

//Pop - 스택에 값을 제거하고 top위치에 값을 반환하는 함수.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return nil
	} else {
		top := len(*s) - 1
		data := (*s)[top] // top 위치에 있는 값을 가져 옴
		*s = (*s)[:top]   // 스택에 마지막 데이터 제거함
		return data
	}
}
