package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	max := 0
	maxi, maxj := len(grid), len(grid[0])
	view := make([][]bool, maxi, maxi)
	for i := 0; i < maxi; i++ {
		view[i] = make([]bool, maxj, maxj)
	}
	for i := 0; i < maxi; i++ {
		for j := 0; j < maxj; j++ {
			if grid[i][j] == 0 || view[i][j] {
				continue
			}
			view[i][j] = true
			open := make([][2]int, 0)
			size := 1
			open = append(open, [2]int{i, j})
			for lo := len(open); lo > 0; lo = len(open) {
				x, y := open[lo-1][0], open[lo-1][1]
				open = open[0 : lo-1]

				if x+1 < maxi && !view[x+1][y] && grid[x+1][y] > 0 {
					view[x+1][y] = true
					open = append(open, [2]int{x + 1, y})
					size++
				}
				if x-1 >= 0 && !view[x-1][y] && grid[x-1][y] > 0 {
					view[x-1][y] = true
					open = append(open, [2]int{x - 1, y})
					size++
				}
				if y+1 < maxj && !view[x][y+1] && grid[x][y+1] > 0 {
					view[x][y+1] = true
					open = append(open, [2]int{x, y + 1})
					size++
				}
				if y-1 >= 0 && !view[x][y-1] && grid[x][y-1] > 0 {
					view[x][y-1] = true
					open = append(open, [2]int{x, y - 1})
					size++
				}
			}
			if size > max {
				max = size
			}
		}
	}
	return max
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	fmt.Println(maxAreaOfIsland(grid))
}
