package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	total_Bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file "+fileName, err)
	}
	total_String := string(total_Bytes)

	rows := strings.Split(total_String, "\n")

	left_list := []int{}
	right_list := []int{}

	for _, row := range rows {
		items := strings.Split(row, "   ")
		if items[0] == "" || items[1] == "" {
			continue
		}

		if left_item, err := strconv.Atoi(items[0]); err != nil {
			log.Fatal("failed to convert left string to int: ", err)
		} else {
			left_list = append(left_list, left_item)
		}

		if right_item, err := strconv.Atoi(items[1]); err != nil {
			log.Fatal("Failed to convert right string to int: ", err)
		} else {
			right_list = append(right_list, right_item)
		}
	}
	sort.Ints(left_list)
	sort.Ints(right_list)
	part_1(left_list, right_list)
	part_2(left_list, right_list)

}

func part_1(left_list []int, right_list []int) {

	count := 0
	for i, _ := range left_list {
		count += int(math.Abs(float64(left_list[i] - right_list[i])))
	}
	fmt.Println(count)
}

func part_2(left_list []int, right_list []int) {
	total := 0
	for _, v := range left_list {
		for _, j := range right_list {
			if v == j {
				total += v
			}
		}
	}
	fmt.Println(total)
}
