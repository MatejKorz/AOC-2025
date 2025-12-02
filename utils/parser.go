package utils

import (
	"bufio"
	"os"
	"strings"
)

func ParseLines(path string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	return ch
}

func ParseDelim(path string, delim string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, delim)
			for _, p := range parts {
				ch <- p
			}
		}
		close(ch)
	}()

	return ch
}
