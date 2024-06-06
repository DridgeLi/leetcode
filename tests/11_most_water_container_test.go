package tests

import (
	"leetcode/solutions"
	"testing"
)

func TestMaxAreaS1(t *testing.T) {
	input := []int{1, 2, 3}
	testMaxArea(t, input, 1)
}

func TestMaxAreaS2(t *testing.T) {
	input := []int{1, 2, 1}
	testMaxArea(t, input, 2)
}

func TestMaxAreaS3(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	testMaxArea(t, input, 49)
}

func testMaxArea(t *testing.T, input []int, expect int) {
	res := solutions.MaxArea(input)
	if res != expect {
		t.Fatalf("fail in MaxArea, input: %v, expect: %v, actual: %v", input, expect, res)
	}
}
