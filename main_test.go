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

func TestGasStations(t *testing.T) {
	type test struct { gas []int
		cost []int
		want int
	}

	tests := []test{
		{
	gas : []int {1,2,3,4,5}
			cost : []int{3,4,5,1,2}
			want:3
		}

	expected := 3
	actual := canCompleteCircuit(gas, cost)
		for _< tc := range tests
	if actual != expected {
		t.Fatalf("Expected: %d, Actual %d", expected, actual)
	}
}
