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
	type test struct {
		gas  []int
		cost []int
		want int
	}

	tests := []test{
		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			gas:  []int{5, 1, 2, 3, 4},
			cost: []int{4, 4, 1, 5, 1},
			want: 4,
		},
		{
			gas:  []int{4, 5, 2, 6, 5, 3},
			cost: []int{3, 2, 7, 3, 2, 9},
			want: -1,
		},
		{
			gas:  []int{4, 5, 3, 1, 4},
			cost: []int{5, 4, 3, 4, 2},
			want: -1,
		},
		{
			gas:  []int{1, 2, 3, 4, 3, 2, 4, 1, 5, 3, 2, 4},
			cost: []int{1, 1, 1, 3, 2, 4, 3, 6, 7, 4, 3, 1},
			want: -1,
			// want: 6690,
		},
	}

	for n, tc := range tests {
		t.Logf("--- Test Case %d", n)
		actual := canCompleteCircuit(tc.gas, tc.cost)
		if tc.want != actual {
			t.Fatalf("Expected: %d, Actual %d", tc.want, actual)
		}
	}
}

func TestTopmostelement(t *testing.T) {
	type test struct {
		elements   []int
		iterations int
		want       int
	}

	tests := []test{
		{
			elements:   []int{5, 2, 2, 4, 0, 6},
			iterations: 4,
			want:       5,
		},
		{
			elements:   []int{2},
			iterations: 1,
			want:       -1,
		},
		{
			elements:   []int{1, 0},
			iterations: 1,
			want:       0,
		},
	}

	for n, tc := range tests {
		t.Logf("--- Test Case %d", n)
		actual := maximumTop(tc.elements, tc.iterations)
		if tc.want != actual {
			t.Fatalf("Expected: %d, Actual %d", tc.want, actual)
		}
	}
}
