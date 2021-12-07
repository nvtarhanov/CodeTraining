package linkedlsit

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	lenght int
}

func (l *LinkedList) AddNode(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.lenght++
}

func (l LinkedList) Print() {
	toPrint := l.head
	for l.lenght != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.lenght--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {

	//If list is empty
	if l.lenght == 0 {
		return
	}

	//If head data equals valuse
	if l.head.Data == value {
		l.head = l.head.Next
		l.lenght--
		return
	}

	previousToDelete := l.head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.lenght--

}
