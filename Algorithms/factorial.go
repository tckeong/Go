package algorithms

import _"fmt"


func Factorial(x int) int {
	if x == 1 {
		return x
	} else {
		x *= factorial_1(x-1)
	}
	return x
}

func Factorial_2(x int) int {
	for i := x-1; i >= 1; i-- {
		x *= i
	}
	return x
}

func factorial_1(x int) int {
	if x == 1 {
		return x
	} else {
		x *= factorial_1(x-1)
	}
	return x
}

func factorial_2(x int) int {
	for i := x-1; i >= 1; i-- {
		x *= i
	}
	return x
}

/*
func main(){
	sum := 0
	fmt.Println("Please enter the sum: ")
	fmt.Scan(&sum)
	x := factorial_1(sum)
	y := factorial_2(sum)
	fmt.Printf("factorial 1 = %v, factorial 2 = %v\n", x, y)
}
*/