package graph

import (
	"fmt"
	common "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_common"
	"testing"
)

func TestGraphDemo(*testing.T) {
	graph := directedGraph(DEMO)
	graph.Print()
	fmt.Println("\nRemove Edge {from=V0, to=V4}")
	graph.RemoveEdge("V0", "V4")
	graph.Print()
	fmt.Println("\nRemove Vertex V0")
	graph.RemoveVertex("V2")
	graph.Print()
}
func TestGraphBreadthFirstSearch(*testing.T) {
	graph := directedGraph(BFS_01)
	graph.breadthFirstSearch("A")
	graph = directedGraph(BFS_02)
	graph.breadthFirstSearch(0)
}
func TestGraphDepthFirstSearch(*testing.T) {
	graph := undirectedGraph(DFS_01)
	graph.depthFirstSearch(0)
	graph = directedGraph(DFS_02)
	graph.depthFirstSearch("d")
}
func TestTopologicalSort(*testing.T) {
	graph := directedGraph(TOPO)
	list := graph.topologicalSort()
	fmt.Printf("%#v\n", list)
}
func TestMinimumSpanningTreePrim(*testing.T) {
	graph := undirectedGraph(MST_01)
	// graph := undirectedGraph(MST_02)
	edgeInfos := graph.prim()
	for edgeInfo := range edgeInfos {
		fmt.Println(edgeInfo.ToString())
	}
}
func TestMinimumSpanningTreeKruskal(*testing.T) {
	graph := undirectedGraph(MST_01)
	edgeInfos := graph.kruskal(common.IntComparator)
	// graph := undirectedGraph(MST_02)
	// edgeInfos := graph.kruskal(common.StringComparator)
	for edgeInfo := range edgeInfos {
		fmt.Println(edgeInfo.ToString())
	}
}
func TestShortestPathDijkstra(*testing.T) {
	// graph := undirectedGraph(SP)
	// graph := undirectedGraph(NEGATIVE_WEIGHT1)
	// sp := graph.dijkstra("A")
	// for v, path := range sp {
	// 	fmt.Printf("%v - %v \n", v.(string), path.ToString())
	// }
	graph := undirectedGraph(NEGATIVE_WEIGHT2)
	sp := graph.dijkstra(0)
	for v, path := range sp {
		fmt.Printf("%v - %v \n", v.(int), path.ToString())
	}

}
func TestShortestPathBellmanFord(*testing.T) {
	graph := undirectedGraph(SP)
	// graph := undirectedGraph(NEGATIVE_WEIGHT1)
	sp := graph.dijkstra("A")
	for v, path := range sp {
		fmt.Printf("%v - %v \n", v.(string), path.ToString())
	}
	//graph := undirectedGraph(NEGATIVE_WEIGHT2)
	//sp := graph.bellmanFord(0)
	//for v, path := range sp {
	//	fmt.Printf("%v - %v \n", v.(int), path.ToString())
	//}
}

//func TestA(*testing.T) {
//	var a1, b1 interface{}
//	a1 = 1
//	b1 = 1
//	fmt.Println(a1 == b1)
//	a1 = "string"
//	b1 = "string"
//	fmt.Println(a1 == b1)
//}

func TestShortestPathFloyd(*testing.T) {
	graph := directedGraph(SP)
	// graph := directedGraph(NEGATIVE_WEIGHT1)
	sp := graph.floyd(common.StringComparator)
	for from, m := range sp {
		fmt.Println(from.(string) + "------------------------")
		for to, path := range m {
			fmt.Printf("%v - %v \n", to.(string), path.ToString())
		}

	}
	//graph := directedGraph(NEGATIVE_WEIGHT2)
	//sp := graph.bellmanFord(0)
	//for v, path := range sp {
	//	fmt.Printf("%v - %v \n", v.(int), path.ToString())
	//}
}

// undirectedGraph 有向图
func directedGraph(data [][]interface{}) (graph Graph) {
	graph = NewListGraph()
	for _, edge := range data {
		length := len(edge)
		if length == 1 {
			graph.AddVertex(edge[0])
		} else if length == 2 {
			graph.AddEdge(edge[0], edge[1])
		} else if length == 3 {
			graph.AddEdgeWithWeight(edge[0], edge[1], NewWeightImpl(edge[2]))
		}
	}
	return
}

// undirectedGraph 无向图
func undirectedGraph(data [][]interface{}) (graph Graph) {
	graph = NewListGraph()
	for _, edge := range data {
		length := len(edge)
		if length == 1 {
			graph.AddVertex(edge[0])
		} else if length == 2 {
			graph.AddEdge(edge[0], edge[1])
			graph.AddEdge(edge[1], edge[0])
		} else if length == 3 {
			graph.AddEdgeWithWeight(edge[0], edge[1], NewWeightImpl(edge[2]))
			graph.AddEdgeWithWeight(edge[1], edge[0], NewWeightImpl(edge[2]))
		}
	}
	return
}

var (
	DEMO = [][]interface{}{
		{"V1", "V0", 9}, {"V1", "V2", 3},
		{"V2", "V0", 2}, {"V2", "V3", 5},
		{"V3", "V4", 1},
		{"V0", "V4", 6},
	}
	BFS_01 = [][]interface{}{
		{"A", "B"}, {"A", "F"},
		{"B", "C"}, {"B", "I"}, {"B", "G"},
		{"C", "I"}, {"C", "D"},
		{"D", "I"}, {"D", "G"}, {"D", "E"}, {"D", "H"},
		{"E", "H"}, {"E", "F"},
		{"F", "G"},
		{"G", "H"},
	}
	BFS_02 = [][]interface{}{
		{0, 1}, {0, 4},
		{1, 2},
		{2, 0}, {2, 4}, {2, 5},
		{3, 1},
		{4, 6}, {4, 7},
		{5, 3}, {5, 7},
		{6, 2}, {6, 7},
	}
	BFS_03 = [][]interface{}{
		{0, 2}, {0, 3},
		{1, 2}, {1, 3}, {1, 6},
		{2, 4},
		{3, 7},
		{4, 6},
		{5, 6},
		{6, 7},
	}
	BFS_04 = [][]interface{}{
		{1, 2}, {1, 3}, {1, 5},
		{2, 0},
		{3, 5},
		{5, 6}, {5, 7},
		{6, 2},
		{7, 6},
	}
	DFS_01 = [][]interface{}{
		{0, 1},
		{1, 3}, {1, 5}, {1, 6}, {1, 2},
		{2, 4},
		{3, 7},
	}
	DFS_02 = [][]interface{}{
		{"a", "b"}, {"a", "e"},
		{"b", "e"},
		{"c", "b"},
		{"d", "a"},
		{"e", "c"}, {"e", "f"},
		{"f", "c"},
	}
	TOPO = [][]interface{}{
		{0, 2},
		{1, 0},
		{2, 5}, {2, 6},
		{3, 1}, {3, 5}, {3, 7},
		{5, 7},
		{6, 4},
		{7, 6},
	}
	NO_WEIGHT2 = [][]interface{}{
		{0, 3},
		{1, 3}, {1, 6},
		{2, 1},
		{3, 5},
		{6, 2}, {6, 5},
		{4, 7},
	}
	NO_WEIGHT3 = [][]interface{}{
		{0, 1}, {0, 2},
		{1, 2}, {1, 5},
		{2, 4}, {2, 5},
		{5, 6}, {7, 6},
		{3},
	}
	MST_01 = [][]interface{}{
		{0, 2, 2}, {0, 4, 7},
		{1, 2, 3}, {1, 5, 1}, {1, 6, 7},
		{2, 4, 4}, {2, 5, 3}, {2, 6, 6},
		{3, 7, 9},
		{4, 6, 8},
		{5, 6, 4}, {5, 7, 5},
	}
	MST_02 = [][]interface{}{
		{"A", "B", 17}, {"A", "F", 1}, {"A", "E", 16},
		{"B", "C", 6}, {"B", "D", 5}, {"B", "F", 11},
		{"C", "D", 10},
		{"D", "E", 4}, {"D", "F", 14},
		{"E", "F", 33},
	}
	WEIGHT3 = [][]interface{}{
		{"广州", "佛山", 100}, {"广州", "珠海", 200}, {"广州", "肇庆", 200},
		{"佛山", "珠海", 50}, {"佛山", "深圳", 150},
		{"肇庆", "珠海", 100}, {"肇庆", "南宁", 150},
		{"珠海", "南宁", 350}, {"珠海", "深圳", 100},
		{"南宁", "香港", 500}, {"南宁", "深圳", 400},
		{"深圳", "香港", 150},
	}
	SP = [][]interface{}{
		{"A", "B", 10}, {"A", "D", 30}, {"A", "E", 100},
		{"B", "C", 50},
		{"C", "E", 10},
		{"D", "C", 20}, {"D", "E", 60},
	}
	BF_SP = [][]interface{}{
		{"A", "B", 10}, {"A", "E", 8},
		{"B", "C", 8}, {"B", "E", -5},
		{"D", "C", 2}, {"D", "F", 6},
		{"E", "D", 7}, {"E", "F", 3},
	}
	WEIGHT5 = [][]interface{}{
		{0, 14, 1}, {0, 4, 8},
		{1, 2, 9},
		{2, 3, 6}, {2, 5, 9},
		{3, 17, 1}, {3, 10, 4},
		{4, 5, 2}, {4, 8, 2},
		{5, 6, 6}, {5, 8, 1}, {5, 9, 4},
		{6, 9, 8},
		{7, 11, 4},
		{8, 9, 1}, {8, 11, 2}, {8, 12, 7},
		{9, 10, 7}, {9, 13, 4},
		{10, 13, 2},
		{11, 12, 7}, {11, 15, 4},
		{12, 13, 2}, {12, 16, 2},
		{13, 16, 7},
		{15, 16, 7}, {15, 17, 7},
		{16, 17, 2},
	}
	NEGATIVE_WEIGHT1 = [][]interface{}{
		{"A", "B", -1}, {"A", "C", 4},
		{"B", "C", 3}, {"B", "D", 2}, {"B", "E", 2},
		{"D", "B", 1}, {"D", "C", 5},
		{"E", "D", -3},
	}
	// NEGATIVE_WEIGHT2 有负权环
	NEGATIVE_WEIGHT2 = [][]interface{}{
		{0, 1, 1},
		{1, 2, 7},
		{1, 0, -2},
	}
)
