package binarysearchthree

// 				100
//			20		 30
//		  10   15   3	5	//Leaves

//If smaller node is on the left and bigger node is on the right it is called
//Binary search three, 20  30  for example

//If each node has 2 or less children - it is called binary three

//INSERT - if we want to insert - we should add a leave, starting to the top
//we are compare inserted value for each node
//for example:

//we want to add 5
//5 is less than 100 so it goes to the left
//5 is smaller than 20 it goe to the left
//5 is smaller than 10 it goes to the left

//Result

// 				100
//			20		 30
//		  10   15   3	5
//		5

//Well three is balansed it is very fast to search(Less than o(n) , it will be o(h)< where h is a height of a three)

//Node represent the component of a binary search tree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//INSERT
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search will take in a key value and return true if a node exists

func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}
