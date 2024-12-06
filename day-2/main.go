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
	fileName := "input.txt"
	total_Bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file "+fileName, err)
	}
	total_String := string(total_Bytes)

	rows := strings.Split(total_String, "\n")

	//m := make(map[int]string)
	s := 0
	for _, r := range rows {
		prev := 0
		desc := true
		broken := false
		for j, num := range strings.Split(r, " ") {
			if num == "" {
				continue
			}

			inum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("Failed to convert num string to int: ", err)
			}
			if j == 0 {
				prev = inum
				continue
			}

			dif := math.Abs(float64(inum - prev))
			if dif > 3 || dif < 1 {
				broken = true
				fmt.Println(fmt.Sprintf("Diff change between %d and %d in row %v", prev, inum, r))
				break
			}

			if j == 1 {
				desc = inum < prev
			}

			if desc != (inum < prev) {
				broken = true
				fmt.Println(fmt.Sprintf("Swapped direction between %d and %d in row %v", prev, inum, r))
				break
			}

			prev = inum
		}
		if broken == false {
			s += 1
		}
	}
	fmt.Println(s)

}
