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
	} else if i%5 == 0 {
		s = s + "Bar"
	} else if i%7 == 0 {
		s = s + "Qix"
	} else {
		return input
	}

	if i == 3 {
		s = s + "Foo"
	} else if i == 5 {
		s = s + "Bar"
	} else if i == 7 {
		s = s + "Qix"
	}

	return s

	// if input == "3" {
	// 	return "Foo"
	// }

	// if input == "5" {
	// 	return "Bar"
	// }

	// if input == "7" {
	// 	return "Qix"
	// } else {
	// 	return input
	// }

	// return "0"
	// return "1"
}
