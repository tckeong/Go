package algorithms

import "fmt"

func Arrange_Int(step int, sum *int, array, result, book []int) {
	size := len(array)
	if step == size {
		fmt.Println(AtoNum(result))
		(*sum)++
		return
	}
	for i := 0; i < size; i++ {
		if book[i] == 0 {
			result[step] = array[i]
			book[i] = 1
			Arrange_Int(step+1, sum, array, result, book)
			book[i] = 0
		}
	}
	return
}

func Arrange_Str(step int, sum *int, array, result []string, book []int) {
	size := len(array)
	if step == size {
		fmt.Println(AtoStr(result))
		(*sum)++
		return
	}
	for i := 0; i < size; i++ {
		if book[i] == 0 {
			result[step] = array[i]
			book[i] = 1
			Arrange_Str(step+1, sum, array, result, book)
			book[i] = 0
		}
	}
	return
}

func AtoNum (array []int) int {
	sum := 0
	for _, v := range array {
		sum = sum*10 + v
	}
	return sum
}

func AtoStr (array []string) string {
	result := ""
	for _, v := range array {
		result += v
	}
	return result
} 

