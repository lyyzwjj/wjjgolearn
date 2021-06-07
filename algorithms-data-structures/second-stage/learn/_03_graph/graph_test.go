package graph

import (
	"fmt"
	"testing"
)

func TestGraph(*testing.T) {
	graph := NewListGraph()
	p1 := 9
	graph.AddEdgeWithWeight(1, 0, &p1)
	p2 := 3
	graph.AddEdgeWithWeight(1, 2, &p2)
	p3 := 2
	graph.AddEdgeWithWeight(2, 0, &p3)
	p4 := 5
	graph.AddEdgeWithWeight(2, 3, &p4)
	p5 := 1
	graph.AddEdgeWithWeight(3, 4, &p5)
	// p6 := 6
	// graph.AddEdgeWithWeight(0, 4, &p6)
	graph.AddEdge(0, 4)
	graph.Print()
	fmt.Println("\nRemove Edge {from=0, to=4}")
	graph.RemoveEdge(0, 4)
	graph.Print()
	fmt.Println("\nRemove Vertex 0")
	graph.RemoveVertex(0)
	graph.Print()
}
