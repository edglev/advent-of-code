package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	// some nums share letters, e.g. fiveight
	numberReplacer = strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	full := 0
	fullConverted := 0

	for scanner.Scan() {
		line := scanner.Text()

		full += findNumbers(line, false)
		fullConverted += findNumbers(line, true)
	}

	fmt.Printf("part 1: %d\npart2: %d", full, fullConverted)

}

func findNumbers(line string, convert bool) int {
	if convert {
		line = replace(line)
	}

	first := string(line[strings.IndexFunc(line, unicode.IsDigit)])
	last := string(line[strings.LastIndexFunc(line, unicode.IsDigit)])

	number, err := strconv.Atoi(first + last)
	if err != nil {
		panic(err)
	}

	return number
}

func replace(s string) string {
	replaced := numberReplacer.Replace(s)
	if s != replaced {
		return replace(replaced)
	}

	return s
}
