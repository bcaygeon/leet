package main

import (
	"testing"
)

func TestPoints(t *testing.T) {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}

	expected := []int{3, 2, 2}
	answers := countPoints(points, queries)

	if len(answers) != len(queries) {
		t.Fatal("Incomplete solution")
	}

	for i, n := range answers {
		if n != expected[i] {
			t.Fatalf("Expected: %d, Actual: %d", expected[i], n)
		}
	}

}
