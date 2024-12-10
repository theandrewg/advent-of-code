package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println(part1("./day_1.txt"))
	fmt.Println(part2("./day_1.txt"))
}

func part1(path string) int {
	total := 0
	left, right := getLocations(path)
	for i := 0; i < len(left); i++ {
		if right[i] >= left[i] {
			total += right[i] - left[i]
			continue
		}
		total += left[i] - right[i]
	}

	return total
}

func part2(path string) int {
	left, right := getLocations(path)
	total := 0
	for _, l := range left {
		c := 0
		for _, r := range right {
			if l == r {
				c++
			}
		}
		total += l * c
	}
    return total
}

func getLocations(path string) (left []int, right []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	for {
		// Example line:88159   51481
		var num int
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		num, err = strconv.Atoi(string(bytes[0:5]))
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, num)

		num, err = strconv.Atoi(string(bytes[8:13]))
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
