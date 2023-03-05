package algorithms

import _"fmt"

func Bucket_sort(array [11]int) (result [11]int) {
	result = [11]int{0}
	for _, val := range array {
		result[val]++;
	}
	return
}

func bucket_sort(array [11]int) (result [11]int) {
	result = [11]int{0}
	for _, val := range array {
		result[val]++;
	}
	return
}

/*
func main() {
	array := [11]int{0}
	sum := 0
	fmt.Println("Please enter the sum: ")
	fmt.Scanln(&sum)
	fmt.Println("Please enter the following array: ")
	for i := 0; i < sum; i++ {
		input := 0
		fmt.Scan(&input)
		array[i] = input
	}
	fmt.Println("Input Done!")
	sorted_array := bucket_sort(array)
	fmt.Println("The sorted array is", sorted_array)
	for i := 0; i < 11; i++ {
		fmt.Printf("score: %v, people: %v\n", i, sorted_array[i])
	}
}*/