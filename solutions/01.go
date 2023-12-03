package solutions

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func getCalibrationForTaskA(line string) (int, error) {
	var first, last rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			d := c
			if first == 0 {
				first = d
			}
			last = d
		}
	}

	calibration, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		return -1, err
	}

	return calibration, nil
}

var digits = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func getCalibrationForTaskB(line string) (int, error) {
	var first, last rune
	var partials []string
	for _, c := range line {
		if unicode.IsDigit(c) {
			// Digit found
			// Clear previous text since we didn't find any matches
			partials = []string{}
			// Assign the first nubmer if we didn't found any until now
			if first == 0 {
				first = c
			}
			// Current found digit will always be the last
			last = c
		} else {
			str := string(c)
			for i, partial := range partials {
				partial += str
				digit, ok := digits[partial]
				if ok {
					// Digit found as text
					// Clear previous text since we didn't find any matches
					partials = []string{}
					// Assign the first nubmer if we didn't found any until now
					if first == 0 {
						first = digit
					}
					// Current found digit will always be the last
					last = digit
					// Stop looking for the rest of the partials, since we already found a digit
					break
				}

				// If partial is not a digit, store it for next iteration
				partials[i] = partial
			}

			// Append current char as a new partial
			partials = append(partials, str)
		}
	}

	strCalibration := string(first) + string(last)
	return strconv.Atoi(strCalibration)
}

func Day01(scanner *bufio.Scanner, task string) (int, error) {
	var cmd func(string) (int, error)
	if task == "a" {
		cmd = getCalibrationForTaskA
	} else if task == "b" {
		cmd = getCalibrationForTaskB
	} else {
		return -1, fmt.Errorf("Unknown taks: %v\n", task)
	}

	var result int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		calibration, err := cmd(line)
		if err != nil {
			return -1, err
		}

		result += calibration
	}

	return result, nil
}
