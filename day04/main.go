package main

import (
	"fmt"

	"github.com/MatejKorz/AOC-2025/utils"
)

func countNeighbours(grid []string, x int, y int) int {
	var cnt int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !((x+j) >= 0 && (x+j) < len(grid[0]) && (y+i) >= 0 && (y+i) < len(grid)) {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			if grid[y+i][x+j] != '.' {
				cnt++
			}
		}
	}
	return cnt
}

func Task1() {
	grid := utils.ParseGrid("./data/day04.txt")

	total := 0
	for y, row := range grid {
		for x, char := range row {
			if char == '.' {
				continue
			}
			cnt := countNeighbours(grid, x, y)
			if cnt < 4 {
				total++
				rowRunes := []rune(grid[y])
				rowRunes[x] = 'X'
				grid[y] = string(rowRunes)
			}
		}
	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	fmt.Println("Total movable rolls:", total)
}

func Task2() {
	grid := utils.ParseGrid("./data/day04.txt")

	total := 0

	for {
		removed := 0
		for y, row := range grid {
			for x, char := range row {
				if char == '.' {
					continue
				}
				cnt := countNeighbours(grid, x, y)
				if cnt < 4 {
					total++
					removed++
					rowRunes := []rune(grid[y])
					rowRunes[x] = '.'
					grid[y] = string(rowRunes)
				}
			}
		}

		if removed == 0 {
			break
		}

	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	fmt.Println("Total movable rolls:", total)
}

func main() {
	Task2()
}
