package hitopia

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

var (
	opening     = []string{"{", "(", "["}
	closing     = []string{"}", ")", "]"}
	closingPair = map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) <= 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func BalancedBracket(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	var exampleStack Stack
	for _, v := range str {
		valStr := string(v)

		switch {
		case slices.Contains[[]string, string](opening, valStr):
			exampleStack.Push(valStr)
		case slices.Contains[[]string, string](closing, valStr):
			if exampleStack.IsEmpty() || (exampleStack[len(exampleStack)-1] != closingPair[valStr]) {
				return "NO"
			} else {
				exampleStack.Pop()
			}
		}
	}

	if len(exampleStack) <= 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func TestBalancedBracket(t *testing.T) {
	t.Run("success, with balanced brackets", func(t *testing.T) {
		res := BalancedBracket("{ [ ( ) ] }")
		assert.Equal(t, "YES", res)
	})

	t.Run("success, with unbalanced brackets", func(t *testing.T) {
		res := BalancedBracket("{ [ ( ] ) }")
		assert.Equal(t, "NO", res)
	})
}
