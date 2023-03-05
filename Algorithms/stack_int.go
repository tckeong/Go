package algorithms

import _"fmt"

type Stack_int struct {
	Array []int
}

func (s Stack_int) Push_int(x int) Stack_int {
	s.Array = append(s.Array, x)
	return s
}

func (s Stack_int) Push_int_slices(x []int) Stack_int {
	s.Array = append(s.Array, x...)
	return s
}

func (s Stack_int) Pop(sum int) ([]int, Stack_int, bool) {
	result := make([]int, 0)
	if sum > len(s.Array) {
		return result, s, false
	}
	for i := 0; i < sum; i++ {
		result = append(result, s.Array[len(s.Array)-1-i])
	}
	s.Array = s.Array[0:len(s.Array)-sum]
	return result, s, true
} 

func (s Stack_int) Head() int {
	return s.Array[0]
}

func (s Stack_int) Tail() int {
	return s.Array[len(s.Array)-1]
}

/*
func main() {
	S := Stack_int{
		Array: []int{1,2,3,4,5},
	}
	fmt.Printf("Head = %v, Tail = %v \n",  S.Head(), S.Tail())
	S = S.Push_int(1)
	fmt.Println("S = ", S.Array)
	S = S.Push_int_slices([]int{1,2,3})
	fmt.Println("S = ", S.Array)
	r, S, _ := S.Pop(5)
	fmt.Println("S = ", S.Array, " r = ", r)

}*/