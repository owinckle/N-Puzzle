package main

// f = g + h
// g = steps from beginning
// h = distance to end
// x and y = coordinates of tile 0
type node struct {
	board   [][]int
	id      string
	f, g, h int
	coord   [2]int
}

var open []*node
var closed []*node

func findNodeIndex() int {
	index := -1
	min := -1
	for i := range open {
		if min < 0 || open[i].f < min {
			min = open[i].f
			index = i
		}
	}
	return index
}
