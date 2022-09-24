package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(braceTest())
	fmt.Println(braceDeepTest())
	var s Stack

	var confirm bool
	str := "{\"t}{1}"
	split := strings.Split(str, "")
	if split[0] != "{" || split[0] == "[" || split[0] == "'" || split[0] == "\"" {
		confirm = false
		fmt.Println(confirm)
	}
	for _, splitStr := range split {
		if splitStr == "{" || splitStr == "[" || splitStr == "'" || splitStr == "\"" {
			s.Push(splitStr)
		}
		if splitStr == "}" || splitStr == "[" || splitStr == "'" || splitStr == "\"" {
			s.Pop()
		}
	}
	if s.IsEmpty() {
		confirm = true
	}
	fmt.Println(confirm)
	fmt.Println(s)
}

func braceTest() bool {
	var s Stack

	var confirm bool
	str := "{t}{1}"
	split := strings.Split(str, "")
	if split[0] != "{" {
		confirm = false
		return confirm
	}
	for _, splitStr := range split {
		if splitStr == "{" {
			s.Push(splitStr)
		}
		if splitStr == "}" {
			s.Pop()
		}
	}
	if s.IsEmpty() {
		confirm = true
	}
	return confirm
}

func braceDeepTest() bool {
	var s Stack

	var confirm bool
	str := "{\"t}{1}"
	split := strings.Split(str, "")
	if split[0] != "{" || split[0] == "[" || split[0] == "'" || split[0] == "\"" {
		confirm = false
		return confirm
	}
	for _, splitStr := range split {
		if splitStr == "{" || splitStr == "[" || splitStr == "'" || splitStr == "\"" {
			s.Push(splitStr)
		}
		if splitStr == "}" || splitStr == "[" || splitStr == "'" || splitStr == "\"" {
			s.Pop()
		}
	}
	if s.IsEmpty() {
		confirm = true
	}
	return confirm
}

type Stack []interface{}

// IsEmpty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push
func (s *Stack) Push(data interface{}) {
	*s = append(*s, data)
}

// Pop
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	} else {
		top := len(*s) - 1
		data := (*s)[top]
		*s = (*s)[:top]
		return data
	}
}
