package main

var right = 1
var down = 2
var left = 3
var up = 4

func generateTable(size int) [][]int {
	dir := right
	tab := make([][]int, size)
	for x := 0; x < size; x++ {
		tab[x] = make([]int, size)
	}
	nbMax := size * size
	i := 0
	j := 0
	for nb := 1; nb < nbMax; nb++ {
		tab[i][j] = nb
		switch dir {
		case right:
			if j+1 < size {
				if tab[i][j+1] == 0 {
					j++
					continue
				}
			}
			dir = down
			i++
		case down:
			if i+1 < size {
				if tab[i+1][j] == 0 {
					i++
					continue
				}
			}
			dir = left
			j--
		case left:
			if j-1 > -1 {
				if tab[i][j-1] == 0 {
					j--
					continue
				}
			}
			dir = up
			i--
		case up:
			if i-1 > -1 {
				if tab[i-1][j] == 0 {
					i--
					continue
				}
			}
			dir = right
			j++
		}
	}
	return tab
}
