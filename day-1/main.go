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
		fmt.Println(items)
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
	count := 0
	for i, _ := range left_list {
		count += int(math.Abs(float64(left_list[i] - right_list[i])))
	}
	fmt.Println(count)
}