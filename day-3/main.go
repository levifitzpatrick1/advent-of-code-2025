package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"strconv"
)

const (
	NA = iota
	MUL
	FIRST_PARENTH
	FIRST_NUM
	COMMA
	SECOND_NUM
)

func main() {
	fileName := "input.txt"
	totalBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file "+fileName, err)
	}

	var s scanner.Scanner

	fset := token.NewFileSet()
	file := fset.AddFile(fileName, fset.Base(), len(totalBytes))

	s.Init(file, totalBytes, nil, scanner.ScanComments)

	state := NA
	num_1 := 0
	num_2 := 0
	num_tot := 0

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)

		switch state {
		case NA:
			if lit == "mul" {
				state = MUL
			}
		case MUL:
			if tok.String() == "(" {
				state = FIRST_PARENTH
			} else {
				state = NA
			}
		case FIRST_PARENTH:
			if tok.String() == "INT" {
				num_1, _ = strconv.Atoi(lit)
				state = FIRST_NUM
			} else {
				state = NA
			}
		case FIRST_NUM:
			if tok.String() == "," {
				state = COMMA
			} else {
				num_1 = 0
				state = NA
			}
		case COMMA:
			if tok.String() == "INT" {
				num_2, _ = strconv.Atoi(lit)
				state = SECOND_NUM
			} else {
				state = NA
			}
		case SECOND_NUM:
			if tok.String() == ")" {
				num_tot += num_1 * num_2
			}
			num_1, num_2 = 0, 0
			state = NA
		}

	}

	println(num_tot)
}
