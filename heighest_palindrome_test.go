package hitopia

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loopRecurssive() {

}

func HeighestPalindrome(input string, k int) string {
	intArr := convertInputToArrInt(input)
	l := 0
	r := len(input) - 1

	// make palindrome
	marker := make([]int, len(intArr))

	palindromed, marker, k := makePalindrome(intArr, marker, k, l, r)

	// maximized digits
	maximized, marker, k := maximizedDigits(palindromed, marker, k, l, r)

	res := convertArrIntToArrString(maximized, 0)

	return res
}

func convertInputToArrInt(str string) []int {
	if len(str) <= 0 {
		return []int{}
	}

	num, _ := strconv.Atoi(string(str[0]))

	return append([]int{num}, convertInputToArrInt(str[1:])...)
}

func convertArrIntToArrString(arrInt []int, index int) string {
	if index >= len(arrInt) {
		return ""
	}

	str := strconv.Itoa(arrInt[index])
	remainindStr := convertArrIntToArrString(arrInt, index+1)

	return str + remainindStr
}

func makePalindrome(intArr, marker []int, k, l, r int) ([]int, []int, int) {
	if l > r {
		return intArr, marker, k
	}

	if intArr[l] != intArr[r] {
		if intArr[l] < intArr[r] {
			intArr[l] = intArr[r]
			marker[l] = 1
		} else {
			intArr[r] = intArr[l]
			marker[r] = 1
		}
		k--
	}

	if k < 0 {
		return []int{-1}, []int{}, k
	}

	return makePalindrome(intArr, marker, k, l+1, r-1)
}

func maximizedDigits(intArr, marker []int, k, l, r int) ([]int, []int, int) {
	if l > r {
		return intArr, marker, k
	}

	if l == r && k >= 1 {
		intArr[l] = 9
		return intArr, marker, k
	}

	if intArr[l] < 9 {
		if (marker[l] == 0 && marker[r] == 0) && k >= 2 {
			intArr[l] = 9
			intArr[r] = 9
			k -= 2
		}

		if (marker[l] == 1 || marker[r] == 1) && k >= 1 {
			intArr[l] = 9
			intArr[r] = 9
			k -= 1
		}
	}

	l++
	r--

	return maximizedDigits(intArr, marker, k, l, r)
}

func TestConvertToString(t *testing.T) {
	res := convertInputToArrInt("3943")

	assert.Equal(t, []int{3, 9, 4, 3}, res)
}

func TestConvertArrIntToArrString(t *testing.T) {
	res := convertArrIntToArrString([]int{9, 9, 2, 2, 9, 9}, 0)

	assert.Equal(t, "992299", res)
}

func TestMakePalindrome(t *testing.T) {
	intArr := []int{0, 9, 2, 2, 8, 2}
	marker := make([]int, len(intArr))

	palindromed, marker, k := makePalindrome(intArr, marker, 3, 0, len(intArr)-1)
	assert.Equal(t, []int{2, 9, 2, 2, 9, 2}, palindromed)
	assert.Equal(t, []int{1, 0, 0, 0, 1, 0}, marker)
	assert.Equal(t, 1, k)
}

func TestMaximizedDigits(t *testing.T) {
	intArr := []int{2, 9, 2, 2, 9, 2}
	marker := []int{1, 0, 0, 0, 1, 0}

	maximized, marker, k := maximizedDigits(intArr, marker, 1, 0, len(intArr)-1)
	assert.Equal(t, []int{9, 9, 2, 2, 9, 9}, maximized)
	assert.Equal(t, []int{1, 0, 0, 0, 1, 0}, marker)
	assert.Equal(t, 0, k)
}

func TestHeighestPalindrome(t *testing.T) {
	res := HeighestPalindrome("092282", 3)
	assert.Equal(t, "992299", res)
}
