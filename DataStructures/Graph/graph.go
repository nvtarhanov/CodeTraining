package graphs

import (
	"fmt"
)

// Representations of graph
//1.Adjacency List

// Each noode holds a list holding the neighbour noods
//Example : if 1 pointing to 2 and 4, i means that this two should be in a adjacency list for 1
//	key:1 	2 4
//	key:2
//	key:3
//	key:4
//	key:5

//2.Adjacency matrix
//We put nodes of the graph as rows and colums
//TO
//		1		2		3		4		5
//FROM	//1				1

//2

//3

//4

//5

//If we want to add node we should add row and column

//Graph structure
type Graph struct {
	vertices []*Vertex
}

//Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

//Add vertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex not added")
		fmt.Println(err.Error())
	} else {

	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

//Add edge

func (g *Graph) AddEdge(from, to int) {

	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", fromVertex, toVertex)
		fmt.Println(err.Error())
		return
	}

	if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", fromVertex, toVertex)
		fmt.Println(err.Error())
		return
	}

	//add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

//getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}
