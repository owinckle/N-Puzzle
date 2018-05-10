package main

// f = g + h
// g = steps from beginning
// h = distance to end
// x and y = coordinates of tile 0
type node struct {
	board   [][]int
	id      string
	g, h, f int
	coord   [2]int
}

var open []*node
var closed []*node
var openMax int = 0
