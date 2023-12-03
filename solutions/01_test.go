package solutions

import "testing"

func TestDay01(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	solutions := []int{
		12,
		38,
		15,
		77,
	}

	for i, line := range input {
		expected := solutions[i]
		actual, err := getCalibrationForTaskA(line)

		if err != nil {
			t.Errorf(err.Error())
		}

		if expected != actual {
			t.Errorf("Expected %d but got %d", expected, actual)
		}
	}
}

func TestDay01TaskB(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"oneight",
	}
	solutions := []int{
		29,
		83,
		13,
		24,
		42,
		14,
		76,
		18,
	}

	for i, line := range input {
		expected := solutions[i]
		actual, err := getCalibrationForTaskB(line)

		if err != nil {
			t.Errorf(err.Error())
		}

		if expected != actual {
			t.Errorf("Expected %d but got %d", expected, actual)
		}
	}
}
