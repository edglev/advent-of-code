package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	full := 0
	fullPower := 0

	for scanner.Scan() {
		line := scanner.Text()
		separated := strings.Split(line, ":")

		game, _ := strconv.Atoi(strings.Split(separated[0], " ")[1])
		sets := strings.Split(separated[1], ";")

		valid := true

		maxBlue := 0
		maxRed := 0
		maxGreen := 0

		for _, setLine := range sets {
			cubes := strings.Split(setLine, ",")

			blue := 14
			red := 12
			green := 13

			for _, cubeLine := range cubes {
				c := strings.Split(strings.TrimSpace(cubeLine), " ")

				num, _ := strconv.Atoi(c[0])

				switch true {
				case strings.Contains(c[1], "blue"):
					blue -= num
					maxBlue = max(num, maxBlue)
				case strings.Contains(c[1], "green"):
					green -= num
					maxGreen = max(num, maxGreen)
				case strings.Contains(c[1], "red"):
					red -= num
					maxRed = max(num, maxRed)
				}
			}

			if blue < 0 || green < 0 || red < 0 {
				valid = false
			}

		}

		if valid {
			full += game
		}
		fullPower += maxBlue * maxGreen * maxRed
	}

	fmt.Printf("part1: %d\n", full)
	fmt.Printf("part2: %d", fullPower)
}
