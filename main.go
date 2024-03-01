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

func canCompleteCircuit(gas []int, cost []int) int { return 0 }
