package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "testinput.txt"
	totalBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file "+fileName, err)
	}
	totalString := string(totalBytes)

	rows := strings.Split(totalString, "\n")
	reattempts := []string{}
	safeCount := 0
	retryCount := 0

	for _, row := range rows {
		row = strings.TrimSpace(row)
		nums := strings.Split(row, " ")
		if len(nums) < 2 {
			continue
		}

		passed := parseGroup(nums)

		if passed {
			safeCount++
		} else {
			fmt.Println(nums)
			reattempts = append(reattempts, row)
			retryCount++
		}
	}

	for _, row := range reattempts {
		row = strings.TrimSpace(row)
		nums := strings.Split(row, " ")
		if len(nums) < 2 {
			continue
		}

		passed := false
		for i := 0; i < len(nums); i++ {
			newNums := remove(nums, i)
			if len(newNums) < 2 {
				continue
			}
			passed = parseGroup(newNums)
			if passed {
				break
			}
		}

		if passed {
			safeCount++
		}
	}

	fmt.Println("Retry Count:", retryCount)
	fmt.Println("Safe Reports with Problem Dampener:", safeCount)
}

func parseGroup(nums []string) bool {
	prev := 0
	desc := true
	for i, num := range nums {
		if num == "" {
			continue
		}
		inum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("Failed to convert num string to int: ", err)
		}
		if i == 0 {
			prev = inum
			continue
		}
		if i == 1 {
			desc = inum < prev
		}
		if inum == prev {
			fmt.Println(fmt.Sprintf("Current num = prev num at num %d in row %v", inum, nums))
			return false
		}

		dif := math.Abs(float64(inum - prev))
		if dif > 3 || dif < 1 {
			fmt.Println(fmt.Sprintf("Diff change between %d and %d in row %v", prev, inum, nums))
			return false
		}

		if desc != (inum < prev) {
			fmt.Println(fmt.Sprintf("Swapped direction between %d and %d in row %v", prev, inum, nums))
			return false
		}

		prev = inum
	}

	return true
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
