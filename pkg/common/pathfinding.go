package common

import (
	"sort"
)

type Dijkstra struct {
	Grid  map[Coordinate2]int
	costs map[Coordinate2]int
	Pos   Coordinate2
}

type DijkstraPath struct {
	cost  int
	pos   Coordinate2
	route []Coordinate2
}

func (dikj *Dijkstra) GetShortestPathCost(goalFn func(coordinate2 Coordinate2) bool, neighboursFn func(coordinate2 Coordinate2) []Coordinate2) int {
	queue := make([]DijkstraPath, 1)
	queue[0] = DijkstraPath{pos: dikj.Pos, cost: 0}
	dikj.costs = make(map[Coordinate2]int)
	dikj.costs[dikj.Pos] = 0
	visited := make([]Coordinate2, 0, 50)
	for len(queue) > 0 {
		route := queue[0]
		cell := route.pos
		newCost := route.cost
		queue = queue[1:]
		if Contains(visited, cell) {
			continue
		}
		visited = append(visited, cell)
		currentCost, found := dikj.costs[cell]
		if !found || currentCost > newCost {
			dikj.costs[cell] = newCost
			if goalFn(cell) {
				return newCost
			}
		}
		for _, neighbour := range neighboursFn(cell) {
			if Contains(visited, neighbour) {
				continue
			}
			queue = append(queue, DijkstraPath{cost: newCost + 1, pos: neighbour})
		}
		sort.Slice(queue, func(i, j int) bool {
			a := queue[i]
			b := queue[j]
			return a.cost < b.cost
		})
	}
	return 0
}
