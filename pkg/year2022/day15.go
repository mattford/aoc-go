package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
)

type Day15 struct{}

type sensor struct {
	Pos         common.Coordinate2
	BeaconPos   common.Coordinate2
	MaxDistance int
}

func (p Day15) PartA(lines []string) any {
	sensors := buildGrid15(lines)
	minX := math.MaxInt
	maxX := math.MinInt
	for _, sensor := range sensors {
		minX = int(math.Min(float64(sensor.Pos.X-sensor.MaxDistance-1), float64(minX)))
		maxX = int(math.Max(float64(sensor.Pos.X+sensor.MaxDistance+1), float64(maxX)))
	}
	total := 0
outer:
	for x := minX; x <= maxX; x++ {
		for _, thisSensor := range sensors {
			thisPos := common.Coordinate2{Y: 2000000, X: x}
			dist := common.Manhattan(thisPos, thisSensor.Pos)
			if thisSensor.BeaconPos != thisPos && dist <= thisSensor.MaxDistance {
				total++
				continue outer
			}
		}
	}
	return total
}

func (p Day15) PartB(lines []string) any {
	sensors := buildGrid15(lines)
	maxCoord := 4000000

	for y := 0; y <= maxCoord; y++ {
	outer:
		for x := 0; x <= maxCoord; x++ {
			thisPos := common.Coordinate2{Y: y, X: x}
			for _, thisSensor := range sensors {
				dist := common.Manhattan(thisPos, thisSensor.Pos)
				if dist <= thisSensor.MaxDistance {
					jumpDist := thisSensor.MaxDistance - int(math.Abs(float64(thisSensor.Pos.Y-thisPos.Y)))
					jumpDist += thisSensor.Pos.X - thisPos.X
					x += jumpDist
					continue outer
				}
			}
			return thisPos.X*4000000 + thisPos.Y
		}
	}
	return 0
}

func buildGrid15(lines []string) []sensor {
	sensors := make([]sensor, 0, len(lines))
	for _, line := range lines {
		var sy, sx, by, bx int
		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensorPos := common.Coordinate2{Y: sy, X: sx}
		beaconPos := common.Coordinate2{Y: by, X: bx}
		maxDistance := common.Manhattan(sensorPos, beaconPos)
		sensors = append(sensors, sensor{Pos: sensorPos, MaxDistance: maxDistance, BeaconPos: beaconPos})
	}
	return sensors
}
