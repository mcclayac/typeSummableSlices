package main

import "fmt"

type SummableSlice []int


func (s SummableSlice) sum() int {

	total := 0
	for _, num := range s {
		total += num
	}
	return total

}


func main() {

	var s SummableSlice = SummableSlice{2,4,6,7,8,7}

	s2 := SummableSlice{2,4,6,7,8,7, 50}

	fmt.Printf("The sum of slice is : %d\n\n", s.sum())
	fmt.Printf("The sum of slice is : %d\n\n", s2.sum())

}
