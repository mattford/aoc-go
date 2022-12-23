package year2022

import (
	"fmt"
	"math"
	"strconv"
)

type Day19 struct{}

type Blueprint struct {
	robots      map[string]map[string]int
	maxRequired map[string]int
}

func (p Day19) PartA(lines []string) any {
	blueprints := parseInput19(lines)
	sum := 0
	for i, blueprint := range blueprints {
		globalBest = 0
		earliestGeode = 0
		visited = make(map[string]int)
		best := findBestApproach(blueprint, 0, 0, 0, 24, 1, 0, 0, 0, 0)
		sum += (i + 1) * best
	}
	return sum
}

func (p Day19) PartB(lines []string) any {
	blueprints := parseInput19(lines)
	product := 1
	for _, blueprint := range blueprints[0:3] {
		globalBest = 0
		earliestGeode = 0
		visited = make(map[string]int)
		best := findBestApproach(blueprint, 0, 0, 0, 32, 1, 0, 0, 0, 0)
		product *= best
	}
	return product
}

var globalBest, earliestGeode int
var visited map[string]int

func findBestApproach(blueprint Blueprint, ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes int) int {
	if geodes > 0 && time > earliestGeode {
		earliestGeode = time
	}
	if time == 0 || (geodes == 0 && time < earliestGeode-3) || globalBest >= geodes+rangeSum(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= blueprint.robots["geode"]["ore"] && obsRobots >= blueprint.robots["geode"]["obsidian"] {
		return rangeSum(geodeRobots, geodeRobots+time-1)
	}
	visited[getStringKey([]int{ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes})] = 1

	oreLimitHit := oreRobots >= blueprint.maxRequired["ore"]
	clayLimitHit := clayRobots >= blueprint.maxRequired["clay"]
	obsLimitHit := obsRobots >= blueprint.maxRequired["obsidian"]
	best := 0

	if !oreLimitHit {
		if _, ok := visited[getStringKey([]int{ore + oreRobots, clay + clayRobots, obs + obsRobots,
			time - 1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes + geodeRobots})]; !ok {
			best = int(math.Max(
				float64(best),
				float64(geodeRobots+findBestApproach(
					blueprint, ore+oreRobots, clay+clayRobots, obs+obsRobots,
					time-1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))))
		}
	}
	if ore >= blueprint.robots["ore"]["ore"] && !oreLimitHit {
		if _, ok := visited[getStringKey([]int{ore - blueprint.robots["ore"]["ore"] + oreRobots, clay + clayRobots, obs + obsRobots,
			time - 1, oreRobots + 1, clayRobots, obsRobots, geodeRobots, geodes + geodeRobots})]; !ok {
			best = int(math.Max(
				float64(best),
				float64(geodeRobots+findBestApproach(
					blueprint, ore-blueprint.robots["ore"]["ore"]+oreRobots, clay+clayRobots, obs+obsRobots,
					time-1, oreRobots+1, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))))
		}
	}
	if ore >= blueprint.robots["clay"]["ore"] && !clayLimitHit {
		if _, ok := visited[getStringKey([]int{ore - blueprint.robots["clay"]["ore"] + oreRobots, clay + clayRobots, obs + obsRobots,
			time - 1, oreRobots, clayRobots + 1, obsRobots, geodeRobots, geodes + geodeRobots})]; !ok {
			best = int(math.Max(
				float64(best), float64(geodeRobots+findBestApproach(
					blueprint, ore-blueprint.robots["clay"]["ore"]+oreRobots, clay+clayRobots, obs+obsRobots,
					time-1, oreRobots, clayRobots+1, obsRobots, geodeRobots, geodes+geodeRobots))))
		}
	}
	if ore >= blueprint.robots["obsidian"]["ore"] && clay >= blueprint.robots["obsidian"]["clay"] && !obsLimitHit {
		if _, ok := visited[getStringKey([]int{ore - blueprint.robots["obsidian"]["ore"] + oreRobots, clay - blueprint.robots["obsidian"]["clay"] + clayRobots, obs + obsRobots,
			time - 1, oreRobots, clayRobots, obsRobots + 1, geodeRobots, geodes + geodeRobots})]; !ok {
			best = int(math.Max(
				float64(best), float64(geodeRobots+findBestApproach(
					blueprint, ore-blueprint.robots["obsidian"]["ore"]+oreRobots, clay-blueprint.robots["obsidian"]["clay"]+clayRobots, obs+obsRobots,
					time-1, oreRobots, clayRobots, obsRobots+1, geodeRobots, geodes+geodeRobots))))
		}
	}
	if ore >= blueprint.robots["geode"]["ore"] && obs >= blueprint.robots["geode"]["obsidian"] {
		if _, ok := visited[getStringKey([]int{ore - blueprint.robots["geode"]["ore"] + oreRobots, clay + clayRobots, obs - blueprint.robots["geode"]["obsidian"] + obsRobots,
			time - 1, oreRobots, clayRobots, obsRobots, geodeRobots + 1, geodes + geodeRobots})]; !ok {
			best = int(math.Max(
				float64(best), float64(geodeRobots+findBestApproach(
					blueprint, ore-blueprint.robots["geode"]["ore"]+oreRobots, clay+clayRobots, obs-blueprint.robots["geode"]["obsidian"]+obsRobots,
					time-1, oreRobots, clayRobots, obsRobots, geodeRobots+1, geodes+geodeRobots))))
		}
	}

	globalBest = int(math.Max(float64(best), float64(globalBest)))
	return best
}

func getStringKey(ints []int) string {
	out := ""
	for _, i := range ints {
		out += strconv.Itoa(i) + "/"
	}
	return out
}

func rangeSum(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func parseInput19(lines []string) []Blueprint {
	out := make([]Blueprint, 0)
	for _, line := range lines {
		var id, oreOreCost, clayOreCost, obsOreCost, obsClayCost, geodeOreCost, geodeObsCost int
		fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreOreCost, &clayOreCost, &obsOreCost, &obsClayCost, &geodeOreCost, &geodeObsCost)

		robots := make(map[string]map[string]int)
		robots["ore"] = map[string]int{"ore": oreOreCost}
		robots["clay"] = map[string]int{"ore": clayOreCost}
		robots["obsidian"] = map[string]int{"ore": obsOreCost, "clay": obsClayCost}
		robots["geode"] = map[string]int{"ore": geodeOreCost, "obsidian": geodeObsCost}
		maxes := make(map[string]int)
		maxes["ore"] = int(math.Max(float64(robots["geode"]["ore"]), math.Max(float64(robots["clay"]["ore"]), float64(robots["obsidian"]["ore"]))))
		maxes["clay"] = robots["obsidian"]["clay"]
		maxes["obsidian"] = robots["geode"]["obsidian"]
		out = append(out, Blueprint{
			robots:      robots,
			maxRequired: maxes,
		})
	}
	return out
}
