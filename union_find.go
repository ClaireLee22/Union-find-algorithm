package main

import "fmt"

type UnionFind struct {
	parent []int
	size   []int
	count  int
}

func newUnionFind(numOfElements int) *UnionFind {
	// makeSet
	parent := make([]int, numOfElements)
	size := make([]int, numOfElements)
	for i := 0; i < numOfElements; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
		count:  numOfElements,
	}
}

// Time: O(logn) | Space: O(1)
func (uf *UnionFind) find(node int) int {
	for node != uf.parent[node] {
		// path compression
		uf.parent[node] = uf.parent[uf.parent[node]]
		node = uf.parent[node]
	}
	return node
}

// Time: O(1) | Space: O(1)
func (uf *UnionFind) union(node1, node2 int) {
	root1 := uf.find(node1)
	root2 := uf.find(node2)

	// already in the same set
	if root1 == root2 {
		return
	}

	if uf.size[root1] > uf.size[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += 1
	} else {
		uf.parent[root1] = root2
		uf.size[root2] += 1
	}

	uf.count -= 1
}

func main() {
	edges := [][]int{
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 7},
		{4, 8},
		{5, 8},
	}

	numOfElements := 9

	uf := newUnionFind(numOfElements)
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		uf.union(node1, node2)
	}
	fmt.Println("number of conntected components: ", uf.count)
}

/*
output:
number of conntected components:  3
*/
