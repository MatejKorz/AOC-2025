package main

import (
	"fmt"
	"strconv"

	"github.com/MatejKorz/AOC-2025/utils"
)

func Task1() {
	lines := utils.ParseLines("./data/day01_1.txt")

	dial := 50
	counter := 0
	for line := range lines {
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			dial = (dial - steps) % 100
		} else {
			dial = (dial + steps) % 100
		}

		if dial == 0 {
			counter++
		}
	}

	fmt.Println("Times at 0:", counter)
}

func Task2() {
	lines := utils.ParseLines("./data/day01_1.txt")

	dial := 50
	counter := 0
	for line := range lines {
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			if line[0] == 'L' {
				dial--
			} else {
				dial++
			}

			switch dial {
			case 100:
				dial = 0
			case -1:
				dial = 99
			}

			if dial == 0 {
				counter++
			}
		}
	}

	fmt.Println("Times at 0:", counter)
}

func main() {
	fmt.Println("Solving task 1")
	Task1()
	fmt.Println("Solving task 2")
	Task2()
}
