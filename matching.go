package main

import (
	"math"
	"slices"
	"strings"
)

func matchClosestString(str string, dataset []string) string {
	var bestScore float64
	scoreboard := make(map[float64]string)

	for _, strInDataset := range dataset {
		similarity := similarityBetween(str, strInDataset)
		bestScore = math.Max(similarity, bestScore)

		if similarity == 1 {
			return strInDataset
		}

		scoreboard[similarity] = strInDataset
	}

	return scoreboard[bestScore]
}

func similarityBetween(s1 string, s2 string) float64 {
	s1Letters, s2Letters := strings.Split(s1, ""), strings.Split(s2, "")
	var checkedLetters []string
	var matched int

	for _, letter := range s2Letters {

		if slices.Contains(s1Letters, letter) && !slices.Contains(checkedLetters, letter) {
			matched += 1
		}

		checkedLetters = append(checkedLetters, letter)
	}

	score := float64(matched) / math.Max(float64(len(s1Letters)), float64(len(s2Letters)))

	return score
}
