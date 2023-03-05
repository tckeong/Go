package algorithms

import _"fmt"

func Quick_sort(array []int) {
	if len(array) > 1 {
		left, right := 0, len(array)-1
		i, j := left, right
		temp := array[left]
		for i != j {
			for array[j] >= temp && i < j {
				j--
			}
			for array[i] <= temp && i < j {
				i++
			}
			array[i], array[j] = array[j], array[i]
		}
		array[left], array[i] = array[i], array[left]
		Quick_sort(array[left:i])
		Quick_sort(array[i+1:len(array)])
	}
}

func quick_sort(array []int) {
	if len(array) > 1 {
		left, right := 0, len(array)-1
		i, j := left, right
		temp := array[left]
		for i != j {
			for array[j] >= temp && i < j {
				j--
			}
			for array[i] <= temp && i < j {
				i++
			}
			array[i], array[j] = array[j], array[i]
		}
		array[left], array[i] = array[i], array[left]
		quick_sort(array[left:i])
		quick_sort(array[i+1:len(array)])
	}
}

/*
func main(){
	array := make([]int, 0)
	sum := 0
	fmt.Println("Please enter the sum: ")
	fmt.Scan(&sum)
	fmt.Println("Please enter the array: ")
	for i := 0; i < sum; i++ {
		temp := 0
		fmt.Scan(&temp)
		array = append(array, temp)
	}
	quick_sort(array)
	fmt.Printf("array = %v\n", array) 
}*/