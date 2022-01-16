package sorts

import "fmt"

func InsertionSort() {
	var numbers = []int{5, 4, 1, 2, 3, 7, 110, 120, 145, 6, -1, 8, 121, 100, 80,73,10009}

	fmt.Println(fmt.Sprintf("%v",sort(numbers)))
}

func sort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {

		key := i
		j := i - 1

		for j >= 0 && numbers[key] < numbers[j] {
			tmp := numbers[j]
			numbers[j] = numbers[key]
			numbers[key] = tmp
			j--
			key = j + 1
		}
	}

	return numbers
}

// insertion sort
// max check
// min check
// range check
// 1,2,3,5,6,7,9,110,120,145
// these are the indices
func bruteForceApproach(numbers []int) {
	min := -1
	max := -1

	for index, value := range numbers {
		if index == 0 {
			min++
			max++
			continue
		}

		if numbers[max] < value {
			max++
			continue
		} else if numbers[min] > value {
			//
			numbers = shiftNumbers(min, index, numbers)
			max++
			numbers[min] = value
			continue
		} else {
			// in b/w range check
			for j := 0; j <= index; {
				if numbers[j] < value {
					j++
					continue
				}

				numbers = shiftNumbers(j, index, numbers)

				numbers[j] = value
				max++
				break
			}
		}

	}

	fmt.Println(fmt.Sprintf("%v", numbers))
}

// 5, 9, 1, 2, 3, 7, 110, 120, 145, 6
func shiftNumbers(rangeFrom int, rangeTo int, numbers []int) []int {

	for i := rangeTo; i > rangeFrom; i-- {
		numbers[i] = numbers[i-1]
	}
	return numbers
}
