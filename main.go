package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func countPoints(points [][]int, queries [][]int) []int {
	answers := make([]int, len(queries))

	for i, q := range queries {
		radiusDistance := math.Pow((float64)(q[2]), 2)
		origin := Point{x: q[0], y: q[1]}
		for _, p := range points {
			dist := origin.distance(Point{p[0], p[1]})
			if dist <= radiusDistance {
				answers[i] = answers[i] + 1
			}
		}
	}

	return answers
}

type Point struct {
	x int
	y int
}

func (p Point) distance(t Point) float64 {
	dx := t.x - p.x
	dy := t.y - p.y
	return math.Pow((float64)(dx), 2) + math.Pow((float64)(dy), 2)
}

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	start := -1
	remainingFuel, totalFuel := 0, 0
	for i := 0; i < l; i++ {
		remainingFuel += gas[i] - cost[i]
		totalFuel += gas[i] - cost[i]
		if remainingFuel < 0 {
			remainingFuel = 0
			start = -1
		} else if start == -1 {
			start = i
		}
	}
	if totalFuel < 0 {
		start = -1
	}
	return start
}
