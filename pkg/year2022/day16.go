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
	routes := findBestRoutes(valveMap, "AA", 0, route{score: 0, valvesOpened: make([]string, 0)})
	return routes[0].score
}

func (p Day16) PartB(lines []string) any {
	valveMap := parseInput16(lines)
	routes := findBestRoutes(valveMap, "AA", 4, route{score: 0, valvesOpened: make([]string, 0)})
	best := 0
	for _, myRoute := range routes {
		for _, elephantRoute := range routes {
			if myRoute.score+elephantRoute.score > best && !routesMatch(myRoute.valvesOpened, elephantRoute.valvesOpened) {
				best = myRoute.score + elephantRoute.score
			}
		}
	}
	return best
}

func routesMatch(a, b []string) bool {
	for _, i := range a {
		if common.Contains(b, i) {
			return true
		}
	}
	return false
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

type route struct {
	valvesOpened []string
	score        int
}

func findBestRoutes(valveMap map[string]valve, pos string, turn int, soFar route) []route {
	routes := make([]route, 0)
	thisValve := valveMap[pos]
	nextMap := removeKey(valveMap, pos)
	for i := range nextMap {
		thatValve := nextMap[i]
		if turn+thisValve.neighbours[i] >= 30 {
			continue
		}
		thisScore := soFar.score
		thisScore += thatValve.flowRate * (30 - (turn + thisValve.neighbours[i]))
		thisValvesOpened := make([]string, len(soFar.valvesOpened))
		copy(thisValvesOpened, soFar.valvesOpened)
		thisValvesOpened = append(thisValvesOpened, i)
		thisRoute := route{
			score:        thisScore,
			valvesOpened: thisValvesOpened,
		}
		if len(nextMap) > 1 {
			thisRoutes := findBestRoutes(nextMap, i, turn+thisValve.neighbours[i], thisRoute)
			routes = append(routes, thisRoutes...)
		}
		routes = append(routes, thisRoute)
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].score > routes[j].score
	})
	return routes
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
