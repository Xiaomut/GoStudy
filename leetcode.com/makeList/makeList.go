package makeList

type Node struct {
	Value int
	Next  *Node
}

func MakeNode(nums []int) *Node {

	if len(nums) == 0 {
		return nil
	}

	res := &Node{
		Value: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &Node{Value: nums[i]}
		temp = temp.Next
	}

	return res
}
