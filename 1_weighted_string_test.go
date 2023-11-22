package hitopia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var alphaWeight = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4,
	"e": 5, "f": 6, "g": 7, "h": 8,
	"i": 9, "j": 10, "k": 11, "l": 12,
	"m": 13, "n": 14, "o": 15, "p": 16,
	"q": 17, "r": 18, "s": 19, "t": 20,
	"u": 21, "v": 22, "w": 23, "x": 24,
	"y": 25, "z": 26,
}

// WeightedString is the function for finding weighted substring
func WeightedString(s string, queries []int) []string {
	var currChar string
	var currSubString string
	var subStrings []string

	for i := 0; i < len(s); i++ {
		currChar = string(s[i])

		if i > 0 && string(s[i]) == string(s[i-1]) {
			currSubString += currChar
		} else {
			currSubString = currChar
		}

		subStrings = append(subStrings, currSubString)
	}

	weightSubStrs := findWeightSubStr(subStrings)

	return checkQueries(weightSubStrs, queries)
}

func findWeightSubStr(subStrings []string) map[string]int {
	weightSubStrs := make(map[string]int)

	for _, subStr := range subStrings {
		weightSubStrs[subStr] = sumWeightAlphabet(subStr)
	}

	return weightSubStrs
}

func sumWeightAlphabet(subStr string) int {
	var res int
	for i := 0; i < len(subStr); i++ {
		if val, ok := alphaWeight[string(subStr[i])]; ok {
			res += val
		}
	}

	return res
}

func checkQueries(weightedStr map[string]int, queries []int) []string {
	var res []string
	var isAvailable bool
	for _, v := range queries {
		for _, weight := range weightedStr {
			if v == weight {
				isAvailable = true
			}
		}

		if isAvailable {
			res = append(res, "yes")
		} else {
			res = append(res, "no")
		}
		isAvailable = false
	}

	return res
}

func TestWeightedString(t *testing.T) {
	expected := []string{"yes", "yes", "yes", "yes", "no", "no"}
	res := WeightedString("abccddde", []int{1, 3, 12, 5, 9, 10})

	assert.Equal(t, expected, res)
}

func TestFindWeightSubStr(t *testing.T) {
	subStr := []string{"a", "b", "bb", "c", "cc", "ccc", "d", "dd", "ddd", "dddd"}
	expected := map[string]int{
		"a": 1, "b": 2, "bb": 4,
		"c": 3, "cc": 6, "ccc": 9,
		"d": 4, "dd": 8, "ddd": 12,
		"dddd": 16,
	}
	res := findWeightSubStr(subStr)

	assert.Equal(t, expected, res)
}

func TestCheckQueries(t *testing.T) {
	weightedStr := map[string]int{
		"a": 1, "b": 2, "bb": 4,
		"c": 3, "cc": 6, "ccc": 9,
		"d": 4, "dd": 8, "ddd": 12,
		"dddd": 16,
	}

	queries := []int{6, 8, 4, 18, 10}

	res := checkQueries(weightedStr, queries)
	assert.Equal(t, []string{"yes", "yes", "yes", "no", "no"}, res)
	assert.NotEqual(t, []string{"no", "yes", "yes", "no", "no"}, res)
}

func TestSumWeightAlphabet(t *testing.T) {
	res := sumWeightAlphabet("bbb")

	assert.Equal(t, 6, res)
}
