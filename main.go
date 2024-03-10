package main

import (
	"fmt"
	"log"
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
	entry := -1
tryagain:
	for i := entry + 1; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			entry = i
			break
		}
		if i == len(gas)-1 {
			entry = -1
		}
	}

	if entry == -1 {
		// no possbile entry point found, no need to try the rest of the circuit
		return entry
	}

	i := entry
	tank := gas[entry]
	var visits int
	log.Printf("Starting at station %d with tank %d\n", entry, tank)
	for visits = len(gas); visits > 0; visits-- {
		tank = tank - cost[i]
		if i == len(gas)-1 {
			i = 0
		} else {
			i++
		}
		if tank < 0 {
			break
		}

		tank = tank + gas[i]
		log.Printf("Moved to station %d with tank %d\n", i, tank)
	}

	if visits != 0 {
		if i != entry {
			goto tryagain
		}
		entry = -1
	}
	return entry
}
