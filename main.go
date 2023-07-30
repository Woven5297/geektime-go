package main

import "fmt"

func SliceDelete[Src any](list []Src, index int) []Src {
	var l []Src

	for i := 0; i < len(list)-1; i++ {
		e := list[i]
		if i < index {
			l = append(l, e)
		} else {
			l = append(l, list[i+1])
		}
	}

	return l
}

func main() {
	list := []string{"a", "b", "c", "d"}
	index := 2
	l2 := []int{1, 2, 3}

	result := SliceDelete(list, index)
	result1 := SliceDelete(l2, 0)
	fmt.Println(result) // Output: [a b d]
	fmt.Println(result1)
}
