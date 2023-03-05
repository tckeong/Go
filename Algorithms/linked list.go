package algorithms

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func Add(n *Node, x int) *Node {
	temp := new(Node)
	(*n).Next = temp
	(*temp).Data = x
	return temp
}

func Del(n *Node, index int) (*Node, bool) {
	var previous_node *Node;
	for i := 0; i < index; i++ {
		previous_node = n
		if (*n).Next != nil {
			n = (*n).Next
		} else {
			return n, false
		}
	}
	(*previous_node).Next = (*n).Next
	return n, true
}

func (n Node) Print(size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("Node_%v = %v\n", i, n.Data)
		if n.Next != nil {
			n = *(n.Next)
		} else {
			break
		}
	}
}

func (n Node) Search_val (val int) ([]int, bool) {
	result := make([]int, 0)
	i := 0
	for ; n.Next != nil; i++ {
		if n.Data == val {
			result = append(result, i)
		}
		n = *(n.Next)
	}
	if n.Data == val {
		result = append(result, i)
	}
	if len(result) == 0 {
		return result, false
	}
	return result, true
}

func (n Node) Search_index(index int) (*Node, bool) {
	var temp *Node
	for i := 0; i < index; i++ {
		if n.Next == nil {
			return new(Node), false			
		} 
		temp = n.Next
		n = *temp
	}
	return temp, true
}

func (n Node) Size() (size int) {
	temp := &n
	for temp != nil {
		size++
		temp = temp.Next
	}
	return
}

/*
func main(){
	N := Node {
		Data: 1,
		Next: nil,
	}
	var temp *Node;
	N_ptr := &N
	for i := 1; i < 5; i++ {
		temp = Add(N_ptr, i)
		N_ptr = temp
	}
	Add(N_ptr, 3)
	N.Print(6)
	fmt.Println("delete function")
	temp, _ = Del(&N, 2)
	fmt.Println("temp = ", (*temp).Data)
	N.Print(5)
	r, _ := N.Search_val(3)
	fmt.Println("r = ", r)
	i, _ := N.Search_index(3)
	fmt.Println("i = ", (*i).Data)
	fmt.Println("size = ", N.Size())
}*/