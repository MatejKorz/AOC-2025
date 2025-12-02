package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MatejKorz/AOC-2025/utils"
)

func getDigit(num int, digit int) int {
	return (num / int(math.Pow10(digit))) % 10
}

func Task1() {
	elems := utils.ParseDelim("./data/day02_1.txt", ",")
	total := 0

	for elem := range elems {
		parts := strings.Split(elem, "-")

		min, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := min; i <= max; i++ {
			str := fmt.Sprintf("%d", i)
			isInvalid := true
			if len(str)%2 != 0 {
				continue
			}
			for j := 0; j < len(str)/2; j++ {
				if str[0+j] != str[len(str)/2+j] {
					isInvalid = false
				}
			}
			if isInvalid {
				total += i
			}
		}
	}

	fmt.Println("Sum of invalid IDs", total)
}

func IsRepetitive(elem string) bool {
	for i := 1; i <= len(elem)/2; i++ {
		subElem := elem[0:i]
		var composed string
		for len(composed) < len(elem) {
			composed += subElem
		}
		if composed == elem {
			return true
		}
	}
	return false
}

func Task2() {
	elems := utils.ParseDelim("./data/day02_1.txt", ",")
	total := 0

	for elem := range elems {
		parts := strings.Split(elem, "-")

		min, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := min; i <= max; i++ {
			str := fmt.Sprintf("%d", i)
			if IsRepetitive(str) {
				//fmt.Println(i)
				total += i
			}
		}
	}

	fmt.Println("Sum of invalid IDs", total)
}

func main() {
	fmt.Println("Solving task 1")
	Task1()
	fmt.Println("Solving task 2")
	Task2()
}
