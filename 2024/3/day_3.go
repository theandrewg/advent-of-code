package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(part1("./day_3.txt"))
}

func part1(path string) int {
	total := 0
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := re.FindAll(bytes, -1)
	for _, match := range matches {
		fmt.Println(string(match))
		re2 := regexp.MustCompile("[0-9]{1,3}")
		nums := re2.FindAll(match, 2)
		a, err := strconv.Atoi(string(nums[0]))
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(string(nums[1]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("a: %d, b: %d\n", a, b)
		total += a * b
	}

	return total
}
