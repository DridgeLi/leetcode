package tests

import (
	"leetcode/problems/solution"
	"testing"
)

func TestMaxAreaCase1(t *testing.T) {
	input := []int{1, 1}
	testMaxArea(t, input, 1)
}
func TestMaxAreaCase2(t *testing.T) {
	input := []int{1, 2, 3}
	testMaxArea(t, input, 2)
}

func TestMaxAreaCase3(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	testMaxArea(t, input, 49)
}

func TestMaxAreaCase4(t *testing.T) {
	input := []int{1, 2, 1}
	testMaxArea(t, input, 2)
}

func testMaxArea(t *testing.T, input []int, expected int) {
	res := solution.MaxArea(input)
	if res != expected {
		t.Fatalf("fail to pass test for MaxArea, input: %v, expected: %v, actual: %v", input, expected, res)
	}
}
