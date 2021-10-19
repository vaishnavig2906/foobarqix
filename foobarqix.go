package main

import (
	"strconv"
)

func Compute(input string) string {
	i, err := strconv.Atoi(input)
	if err != nil {
		return "error"
	}

	s := ""
	if i%3 == 0 {
		s = s + "Foo"
	}
	if i%5 == 0 {
		s = s + "Bar"
	}
	if i%7 == 0 {
		s = s + "Qix"
	}

	for j := 0; j < len(input); j++ {
		if input[j] == '3' {
			s = s + "Foo"
		} else if input[j] == '5' {
			s = s + "Bar"
		} else if input[j] == '7' {
			s = s + "Qix"
		}
	}

	if len(s) == 0 {
		return input
	}
	return s
}
