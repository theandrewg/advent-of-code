package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1("./day_2.txt"))
	fmt.Println(part2("./day_2.txt"))
}

func part2(path string) (count int) {
	reports := getReports(path)
	for _, report := range reports {
        // try with all values
		if validReport(report, 1) {
			count++
			continue
		}
        // try with the first level removed
		if validReport(report[1:], 0) {
			count++
			continue
		}
        // try with the second level removed
		if validReport(append(report[:1], report[2:]...), 0) {
			count++
			continue
		}
	}
	return count
}

func part1(path string) (count int) {
	count = 0
	reports := getReports(path)
	for _, report := range reports {
		if validReport(report, 0) {
			count++
		}
	}
	return count
}

func validReport(report []int, dampener int) bool {
	prev := report[0]
	bad := 0
	asc := true

	if report[1] < report[0] {
		asc = false
	}

	for i := 1; i < len(report); i++ {
		level := report[i]
		if asc {
			// if ascending, make sure the current level is between 1 and 3 greater than the previous level
			if !(level <= prev+3 && level >= prev+1) {
				bad++
				if bad > dampener {
                    return false
				}
				continue
			}
		} else {
			// if descending, make sure the current level is between 1 and 3 less than the previous level
			if !(level <= prev-1 && level >= prev-3) {
				bad++
				if bad > dampener {
                    return false
				}
				continue
			}
		}
		// if we make it here, the level must have been valid
		prev = level
	}

	if bad > dampener {
		return false
	}

	return true
}

func getReports(path string) (reports [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		report := []int{}
		for _, token := range strings.SplitN(line, " ", len(line)) {
			token = strings.TrimSuffix(token, "\n")
			level, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports
}
