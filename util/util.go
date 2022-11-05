package util

import (
	"golang.org/x/exp/slices"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PickRandomNumbers(low, high, n int) []int {
	// assert low <= high
	// assert high - low >= n
	return getNumbersInRange(low, high+1)[:n]
}

func getNumbersInRange(min, max int) []int {
	numbers := rand.Perm(max - min)
	for i, _ := range numbers {
		numbers[i] += min
	}
	return numbers
}

func IntArrayToString(a []int, sepOptional ...string) string {
	sep := ", "
	if len(sepOptional) > 0 {
		sep = sepOptional[0]
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}

func ParseBool(s string) bool {
	trueValues := []string{"t", "true", "T", "TRUE", "1"}
	return slices.Contains(trueValues, s)
}
