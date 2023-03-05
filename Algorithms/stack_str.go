package algorithms

type Stack_str struct {
	Array []string
}

func (s Stack_str) Push_str(x string) Stack_str {
	s.Array = append(s.Array, x)
	return s
}

func (s Stack_str) Push_str_slices(x []string) Stack_str {
	s.Array = append(s.Array, x...)
	return s
}

func (s Stack_str) Pop(sum int) ([]string, Stack_str, bool) {
	result := make([]string, 0)
	if sum > len(s.Array) {
		return result, s, false
	}
	for i := 0; i < sum; i++ {
		result = append(result, s.Array[len(s.Array)-1-i])
	}
	s.Array = s.Array[0:len(s.Array)-sum]
	return result, s, true
} 

func (s Stack_str) Head() string {
	return s.Array[0]
}

func (s Stack_str) Tail() string {
	return s.Array[len(s.Array)-1]
}