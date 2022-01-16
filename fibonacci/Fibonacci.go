package main

func PrintFibonacci(number int) int {
	if number < 2 {
		return 1
	}

	return PrintFibonacci(number-1) + PrintFibonacci(number-2)
}
