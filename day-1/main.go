package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ElfCalorieParseAndSum("puzzle-data.txt")
}

func ElfCalorieParseAndSum(path string) []int {
	content, err := ioutil.ReadFile(path)

	// toss an error & kill process
	// if the data cannot be read
	if err != nil {
		log.Fatal("could not read data")
	}

	// split at every line break
	valArr := strings.Split(string(content), "\n")

	// slush values as we parse
	var sumsArr []int
	tempSum := 0

	for i := 0; i < len(valArr); i++ {
		if valArr[i] == "" {
			// if the value is a line separation
			// then append our new sum for one elf
			// and reset our summator
			sumsArr = append(sumsArr, tempSum)
			tempSum = 0
		} else {
			// coerce string to int
			newInt, err := strconv.Atoi(valArr[i])

			// toss an error & kill process
			// if string cannot convert to integer
			if err != nil {
				log.Fatal("could not convert string to int")
			}

			// add the new value to our current summator
			tempSum += newInt
		}
	}

	// sort the array of elf calorie sums
	sort.Ints(sumsArr)

	// reverse the array
	for i, j := 0, len(sumsArr)-1; i < j; i, j = i+1, j-1 {
		sumsArr[i], sumsArr[j] = sumsArr[j], sumsArr[i]
	}

	return sumsArr
}
