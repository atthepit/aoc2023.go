package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Cubes map[string]int
}

const (
	gameDivider   = ": "
	subsetDivider = "; "
	drawsDivider  = ", "
	divider       = " "
)

func getGameIdForTaskA(line string) (int, error) {
	gamesplit := strings.Split(line, gameDivider)
	game, err := strconv.Atoi(strings.Split(gamesplit[0], " ")[1])
	if err != nil {
		return -1, err
	}

	return game, nil
}

func getMaxCubesForTaskA(line string) (map[string]int, error) {
	result := map[string]int{}

	games := strings.Split(strings.Split(line, gameDivider)[1], subsetDivider)
	for _, game := range games {
		draws := strings.Split(game, drawsDivider)
		for _, draw := range draws {
			parts := strings.Split(draw, divider)
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			color := parts[1]

			if current, ok := result[color]; ok && current > count {
				result[color] = current
			} else {
				result[color] = count
			}
		}
	}

	return result, nil
}

func parseGameForTaskA(line string) (Game, error) {
	game, err := getGameIdForTaskA(line)
	if err != nil {
		return Game{}, err
	}

	cubes, err := getMaxCubesForTaskA(line)
	if err != nil {
		return Game{}, err
	}

	return Game{
		Id:    game,
		Cubes: cubes,
	}, nil
}

func isGamePossibleForTaskA(game Game, condition map[string]int) bool {
	cubes := game.Cubes
	for key, expectedCubes := range condition {
		if gameCubes, hasCubes := cubes[key]; !hasCubes || gameCubes > expectedCubes {
			return false
		}
	}
	return true
}

func Day02(scanner *bufio.Scanner, task string) (int, error) {
	result := 0
	condition := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		game, err := parseGameForTaskA(line)
		if err != nil {
			return -1, err
		}

		if isGamePossibleForTaskA(game, condition) {
			result += game.Id
		}
	}

	return result, nil
}
