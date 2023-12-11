package main

import (
	"strings"
)

func d01Parseline(input string) (uint, error) {

	digits := map[string]int{
		"zero": 0, "0": 0,
		"one": 1, "1": 1,
		"two": 2, "2": 2,
		"three": 3, "3": 3,
		"four": 4, "4": 4,
		"five": 5, "5": 5,
		"six": 6, "6": 6,
		"seven": 7, "7": 7,
		"eight": 8, "8": 8,
		"nine": 9, "9": 9,
	}

	first := -1
	firstIndex := len(input) + 1
	last := -1
	lastIndex := -1

	for s, v := range digits {
		f := strings.Index(input, s)
		if f >= 0 && f < firstIndex {
			first = v
			firstIndex = f
		}
		l := strings.LastIndex(input, s)
		if l >= 0 && l > lastIndex {
			last = v
			lastIndex = l
		}

	}

	if first < 0 || last < 0 {
		return 0, nil // errors.New("Input too short")
	}

	var value uint = uint(first*10 + last)

	return value, nil
}

func Day01(input []string) (uint, error) {
	var value uint = 0
	for _, i := range input {
		v, _ := d01Parseline(i)
		value += v
	}
	return value, nil
}
