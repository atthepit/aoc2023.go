package solutions

import (
	"testing"
)

func TestGetGameIdForTaskA(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	solution := []int{
		1,
		2,
		3,
		4,
		5,
	}

	for i, x := range input {
		expected := solution[i]
		actual, err := getGameIdForTaskA(x)

		if err != nil {
			t.Error(err)
		}

		if actual != expected {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}

func TestGetMaxCubesForTaskA(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	solution := []map[string]int{
		{"blue": 6, "red": 4, "green": 2},
		{"blue": 4, "red": 1, "green": 3},
		{"blue": 6, "red": 20, "green": 13},
		{"blue": 15, "red": 14, "green": 3},
		{"blue": 2, "red": 6, "green": 3},
	}

	for i, x := range input {
		expected := solution[i]
		actual, err := getMaxCubesForTaskA(x)

		if err != nil {
			t.Error(err)
		}

		if actual["blue"] != expected["blue"] {
			t.Errorf("Expected blue %v but got %v", expected["blue"], actual["blue"])
		}
		if actual["red"] != expected["red"] {
			t.Errorf("Expected red %v but got %v", expected["red"], actual["red"])
		}
		if actual["green"] != expected["green"] {
			t.Errorf("Expected green %v but got %v", expected["green"], actual["green"])
		}
	}
}

func TestGetGameForTaskA(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	solution := []Game{
		{Id: 1, Cubes: map[string]int{"blue": 6, "red": 4, "green": 2}},
		{Id: 2, Cubes: map[string]int{"blue": 4, "red": 1, "green": 3}},
		{Id: 3, Cubes: map[string]int{"blue": 6, "red": 20, "green": 13}},
		{Id: 4, Cubes: map[string]int{"blue": 15, "red": 14, "green": 3}},
		{Id: 5, Cubes: map[string]int{"blue": 2, "red": 6, "green": 3}},
	}

	for i, x := range input {
		expected := solution[i]
		actual, err := parseGameForTaskA(x)

		if err != nil {
			t.Error(err)
		}

		if actual.Id != expected.Id {
			t.Errorf("expected.Id %v but got %v", expected.Id, actual.Id)
		}
		if actual.Cubes["blue"] != expected.Cubes["blue"] {
			t.Errorf("expected.Cubes blue %v but got %v", expected.Cubes["blue"], actual.Cubes["blue"])
		}
		if actual.Cubes["red"] != expected.Cubes["red"] {
			t.Errorf("expected.Cubes red %v but got %v", expected.Cubes["red"], actual.Cubes["red"])
		}
		if actual.Cubes["green"] != expected.Cubes["green"] {
			t.Errorf("expected.Cubes green %v but got %v", expected.Cubes["green"], actual.Cubes["green"])
		}
	}
}

func TestGetGameIsPossibleForTaskA(t *testing.T) {
	condition := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	input := []Game{
		{Id: 1, Cubes: map[string]int{"blue": 9, "red": 5, "green": 4}},
		{Id: 2, Cubes: map[string]int{"blue": 6, "red": 1, "green": 6}},
		{Id: 3, Cubes: map[string]int{"blue": 11, "red": 25, "green": 26}},
		{Id: 4, Cubes: map[string]int{"blue": 21, "red": 23, "green": 7}},
		{Id: 5, Cubes: map[string]int{"blue": 3, "red": 7, "green": 5}},
	}
	solution := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	for i, x := range input {
		expected := solution[i]
		actual := isGamePossibleForTaskA(x, condition)

		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}
