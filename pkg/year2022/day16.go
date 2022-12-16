package year2022

import (
	"aocgen/pkg/common"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day16 struct{}

type valve struct {
	flowRate   int
	neighbours map[string]int
}

func (p Day16) PartA(lines []string) any {
	valveMap := parseInput16(lines)
	return findBestRoute(valveMap, "AA", 0, 0)
}

func (p Day16) PartB(lines []string) any {
	valveMap := parseInput16(lines)
	return findBestRouteWithElephant(valveMap, "AA", "AA", 0, 4, 4)
}

func parseInput16(lines []string) map[string]valve {
	valveMap := make(map[string]valve)
	expr := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ([A-Z, ]+)")
	for _, line := range lines {
		matches := expr.FindStringSubmatch(line)
		thisId := matches[1]
		flowRate, _ := strconv.Atoi(matches[2])
		neighbours := make(map[string]int)
		for _, v := range strings.Split(matches[3], ", ") {
			neighbours[v] = 1
		}
		valveMap[thisId] = valve{
			flowRate:   flowRate,
			neighbours: neighbours,
		}
	}
	weightedValveMap := make(map[string]valve)
	for me, v := range valveMap {
		if v.flowRate == 0 && me != "AA" {
			continue
		}
		ns := make(map[string]int)
		for n := range valveMap {
			if valveMap[n].flowRate > 0 {
				ns[n] = navigateTunnels(valveMap, me, n) + 1
			}
		}
		weightedValveMap[me] = valve{flowRate: v.flowRate, neighbours: ns}
	}
	return weightedValveMap
}

type valvePath struct {
	pos   string
	cost  int
	score int
}

func navigateTunnels(valveMap map[string]valve, from string, to string) int {
	queue := make([]valvePath, 1)
	queue[0] = valvePath{pos: from, cost: 0}
	visited := make([]string, 0)
	for len(queue) > 0 {
		route := queue[0]
		cost := route.cost
		queue = queue[1:]
		if common.Contains(visited, route.pos) {
			continue
		}
		visited = append(visited, route.pos)
		if route.pos == to {
			return cost
		}
		v := valveMap[route.pos]
		for nPos, nCost := range v.neighbours {
			if !common.Contains(visited, nPos) {
				queue = append(queue, valvePath{cost: cost + nCost, pos: nPos})
			}
		}
		sort.Slice(queue, func(i, j int) bool {
			a := queue[i]
			b := queue[j]
			return a.cost < b.cost
		})
	}
	return 0
}

func findBestRouteWithElephant(valveMap map[string]valve, pos string, elephantPos string, score int, turn int, elephantTurn int) int {
	thisValve := valveMap[pos]
	elephantValve := valveMap[elephantPos]
	nextMap := removeKey(valveMap, pos)
	nextMap = removeKey(nextMap, elephantPos)
	best := score
	for i := range nextMap {
		for i2 := range removeKey(nextMap, i) {
			thatValve := nextMap[i]
			thatElephantValve := nextMap[i2]
			thisScore := score
			nextMeTurn := turn + thisValve.neighbours[i]
			nextElephantTurn := elephantTurn + elephantValve.neighbours[i2]
			if nextMeTurn < 30 {
				thisScore += thatValve.flowRate * (30 - (turn + thisValve.neighbours[i]))
			}
			if nextElephantTurn < 30 {
				thisScore += thatElephantValve.flowRate * (30 - (elephantTurn + elephantValve.neighbours[i2]))
			}
			if nextMeTurn >= 30 && nextElephantTurn >= 30 {
				continue
			}
			if len(nextMap) > 1 {
				thisScore = findBestRouteWithElephant(nextMap, i, i2, thisScore, turn+thisValve.neighbours[i], elephantTurn+elephantValve.neighbours[i2])
			}
			if thisScore > best {
				best = thisScore
			}
		}
	}
	return best
}

func findBestRoute(valveMap map[string]valve, pos string, score int, turn int) int {
	thisValve := valveMap[pos]
	nextMap := removeKey(valveMap, pos)
	best := score
	for i := range nextMap {
		thatValve := nextMap[i]
		if turn+thisValve.neighbours[i] >= 30 {
			continue
		}
		thisScore := score
		thisScore += thatValve.flowRate * (30 - (turn + thisValve.neighbours[i]))
		if len(nextMap) > 1 {
			thisScore = findBestRoute(nextMap, i, thisScore, turn+thisValve.neighbours[i])
		}
		if thisScore > best {
			best = thisScore
		}
	}
	return best
}

func removeKey(m map[string]valve, k string) map[string]valve {
	m2 := make(map[string]valve)
	for i := range m {
		if i != k {
			m2[i] = m[i]
		}
	}
	return m2
}
