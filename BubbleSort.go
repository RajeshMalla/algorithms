package main

import "fmt"

func main() {

	var numbers = []int{5, 4, 1, 2, 3, 7, 110, 120, 145, 6, -1, 8, 121, 100, 80, 73, 10009, -2}

	// if i am greater then i will go forward

	for i := 0; i < len(numbers); i++ {

		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				tmp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = tmp
			}
		}
	}

	fmt.Println(fmt.Sprintf("%v", numbers))
}
