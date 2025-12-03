package main

import (
	"fmt"
	"math"

	"github.com/MatejKorz/AOC-2025/utils"
)

func Task1() {
	lines := utils.ParseLines("./data/day03_1.txt")

	total := 0
	for line := range lines {
		fmt.Println(line)

		fst := '0'
		last := '0'

		runes := []rune(line)
		for i, char := range line {
			if int(fst) < int(char) && i != len(line)-1 {
				fst = char
				last = runes[i+1]
			} else if int(last) < int(char) {
				last = char
			}
		}

		num := int(fst-'0')*10 + int(last-'0')
		fmt.Println(num)
		total += num
	}
	fmt.Println("Total capacity:", total)

}

func Task2() {
	lines := utils.ParseLines("./data/day03_1.txt")

	total := 0
	for line := range lines {
		fmt.Println(line)

		// last elem = highest order
		var digits [12]rune
		for i := range digits {
			digits[i] = '0'
		}

		index := 0
		runes := []rune(line)
		for j := len(digits) - 1; j >= 0; j-- {
			for i := index; i < len(line)-j; i++ {
				if int(digits[j]) < int(runes[i]) {
					digits[j] = runes[i]
					index = i + 1
				}
			}
		}
		num := 0
		for i, digit := range digits {
			num += int(digit-'0') * int(math.Pow10(i))
		}
		fmt.Println(num)
		total += num
	}
	fmt.Println("Total capacity:", total)
}

func main() {
	Task2()
}
