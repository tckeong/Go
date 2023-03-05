package algorithms

type Queue_str struct {
	Array []string
}

func (q Queue_str) Enqueue_str(x string) Queue_str {
	q.Array = append(q.Array, x)
	return q
}

func (q Queue_str) Enqueue_str_slice(x []string) Queue_str {
	q.Array = append(q.Array, x...)
	return q
}

func (q Queue_str) Dequeue(sum int) ([]string, Queue_str, bool) {
	result := make([]string, 0)
	if sum > len(q.Array) {
		return result, q, false
	}
	for i := 0; i < sum; i++ {
		result = append(result, q.Array[i])
	}
	q.Array = q.Array[sum:len(q.Array)]
	return result, q, true
}

func (q Queue_str) Head() string {
	return q.Array[0]
}

func (q Queue_str) Tail() string {
	return q.Array[len(q.Array)-1]
}