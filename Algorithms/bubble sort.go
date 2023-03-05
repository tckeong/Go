package algorithms

import _"fmt"

func Bubble_sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j+1 < len(array); j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func bubble_sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j+1 < len(array); j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

/*
func main() {
	array := make([]int, 0)
	input := 0
	fmt.Println("Please enter the sum: ")
	fmt.Scanln(&input);
	fmt.Println("Please enter the array: ")
	for i := 0; i < input; i++ {
		temp := 0
		fmt.Scan(&temp)
		array = append(array, temp)
	}
	_ = bubble_sort(array)
	fmt.Println("The sorted array is ", array)
}*/