package fibonacci

import (
	"algorithms/fibonacci"
	"fmt"
	"testing"
)

func TestFibonacciSeries(t *testing.T) {
	val := fibonacci.PrintFibonacci(5)
	if val < 0 {
		t.Errorf("This is failed")
	}
	fmt.Println(val)
}