package main

import (
	binarysearchthree "Algoritms/DataStructures/BinarySearchThree"
	"fmt"
)

//Graph structure

// type Graph struct {
// 	vertices []*Vertex
// }

// //Vertex structure
// type Vertex struct {
// 	key      int
// 	adjacent []*Vertex
// }

// //Add vertex
// func (g *Graph) AddVertex(k int) {
// 	g.vertices = append(g.vertices, &Vertex{key: k})
// }

// //Add Edge

// func (g *Graph) AddEdge(from, to int) {

// }

// func (g *Graph) Print() {
// 	for _, v := range g.vertices {
// 		fmt.Printf("\nVertex %v :", v.key)
// 		for _, v := range v.adjacent {
// 			fmt.Printf(" %v ", v.key)
// 		}
// 	}

// 	fmt.Println()
// }

func main() {
	//HeapExample

	// buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	// for _, value := range buildHeap {
	// 	m.Insert(value)
	// }

	// fmt.Println(m)

	// for i := 0; i < 5; i++ {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }

	//Linked list example
	// myList := linkedlsit.LinkedList{}
	// node1 := &linkedlsit.Node{Data: 48}
	// node2 := &linkedlsit.Node{Data: 18}
	// node3 := &linkedlsit.Node{Data: 16}
	// node4 := &linkedlsit.Node{Data: 12}
	// node5 := &linkedlsit.Node{Data: 10}

	// myList.AddNode(node1)
	// myList.AddNode(node2)
	// myList.AddNode(node3)
	// myList.AddNode(node4)
	// myList.AddNode(node5)

	// myList.Print()

	// myList.DeleteWithValue(16)
	// myList.Print()

	//Binary search three

	tree := &binarysearchthree.Node{Key: 100}
	fmt.Println(tree)

}
