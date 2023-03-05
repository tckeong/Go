package algorithms

import _"fmt"

type Queue_int struct {
	Array []int
}

func (q Queue_int) Enqueue_int(x int) Queue_int {
	q.Array = append(q.Array, x)
	return q
}

func (q Queue_int) Enqueue_int_slice(x []int) Queue_int {
	q.Array = append(q.Array, x...)
	return q
}

func (q Queue_int) Dequeue(sum int) ([]int, Queue_int, bool) {
	result := make([]int, 0)
	if sum > len(q.Array) {
		return result, q, false
	}
	for i := 0; i < sum; i++ {
		result = append(result, q.Array[i])
	}
	q.Array = q.Array[sum:len(q.Array)]
	return result, q, true
}

func (q Queue_int) Head() int {
	return q.Array[0]
}

func (q Queue_int) Tail() int {
	return q.Array[len(q.Array)-1]
}

/*
func main() {
	q := Queue_int {
		Array: []int{1,2,3,4,5},
	}
	fmt.Printf("head = %v, tail = %v\n", q.Head(), q.Tail())
	fmt.Println("queue = ", q.Array)
	r, q, _ := q.Dequeue(2)
	fmt.Println("queue = ", q.Array, r)
	q = q.Enqueue_int(10)
	fmt.Println("queue = ", q.Array)
	q = q.Enqueue_int_slice([]int{3,4,5})
	fmt.Println("queue = ", q.Array)
	fmt.Printf("Head = %v, tail = %v\n", q.Head(), q.Tail())
}*/