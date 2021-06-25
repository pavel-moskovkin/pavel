package pipeline

import (
	"fmt"
	"math"
)

var (
	i, distanceHouse, distanceToNextHouse uint
)

type coords struct {
	x, y uint
}

func newCoords(x, y uint) *coords {
	return &coords{
		x: x,
		y: y,
	}
}

type pipe struct {
	start, finish coords
}

type pipeline struct {
	pipes   []pipe
	houses  []uint
	maxDist uint
}

// As it is not a part of a task description, assuming that the road is located from south to north,
// and 'houses' is a list of distances starting from the southernmost point.
// The solution is based on a X/Y coordinates scale with the dots representing gas pipeline location.
func calculateLocation(houses []uint) (result pipeline, totalLength int) {
	pipes := make([]pipe, 0)
	for i = 0; i < uint(len(houses)-1); i++ {

		distanceHouse = houses[i]
		distanceToNextHouse = houses[i+1]

		appendSimplePipeAlongsideRoad := func(distanceFromRoad uint) {
			pipes = append(pipes, pipe{
				start:  *newCoords(distanceFromRoad, i),
				finish: *newCoords(distanceFromRoad, i+1),
			})
		}

		// 1) if equal distance
		if distanceHouse == distanceToNextHouse {
			appendSimplePipeAlongsideRoad(distanceHouse)
			continue
		}

		var currentHouseIsCloserToRoad = false
		diff := int(distanceHouse) - int(distanceToNextHouse)
		if diff < 0 {
			currentHouseIsCloserToRoad = true
		}

		// 2) if current house is closer to road
		if currentHouseIsCloserToRoad {
			appendSimplePipeAlongsideRoad(distanceHouse)

			pipes = append(pipes, pipe{
				start:  *newCoords(distanceHouse, i+1),
				finish: *newCoords(distanceToNextHouse, i+1),
			})
			continue
		}

		// 3) if current house is further from the road than the next house
		{
			if i == 0 {
				pipes = append(pipes, pipe{
					start:  *newCoords(distanceHouse, i),
					finish: *newCoords(distanceToNextHouse, i),
				})
				appendSimplePipeAlongsideRoad(distanceToNextHouse)
				continue
			}

			// 3.1
			// if (houses[i-1] <= houses[i+1]) {
			//   ~~~ using existing pipe, no need to build a pipe perpendicular (Y) to the road ~~~
			// }

			// 3.2 else, append extra length to existing horizontal (X) pipe
			distanceToPrevHouse := houses[i-1]
			if distanceToPrevHouse > distanceToNextHouse {
				pipes = append(pipes, pipe{
					start:  *newCoords(distanceHouse, i), // distanceToPrevHouse
					finish: *newCoords(distanceToNextHouse, i),
				})
			}

			appendSimplePipeAlongsideRoad(distanceToNextHouse)
		}

	}
	result.pipes = pipes
	result.houses = houses

	totalLength = result.calculateCost()
	return result, totalLength
}

func (pl *pipeline) calculateCost() (totalLength int) {
	var l int
	for _, p := range pl.pipes {
		if p.start.x == p.finish.x {
			l = int(math.Abs(float64(int(p.finish.y) - int(p.start.y))))
		} else {
			l = int(math.Abs(float64(int(p.finish.x) - int(p.start.x))))
		}
		totalLength += l
	}
	return totalLength
}

func (pl *pipeline) printPipeline() {
	matrix := prepareMatrix(pl)

	for _, pipe := range pl.pipes {
		matrix[pipe.start.y][pipe.start.x] = "."
		matrix[pipe.finish.y][pipe.finish.x] = "."
	}

	fmt.Print("\n")
	for j := len(pl.houses) - 1; j >= 0; j-- {
		fmt.Printf("%v    |      |      |   %v house[%d] dist=%v \n", j, matrix[uint(j)], j, pl.houses[j])
	}
	xLine := make([]int, pl.maxDist+1)
	for i := 0; i < int(pl.maxDist+1); i++ {
		xLine[i] = i
	}
	fmt.Printf("%v  | ^ R O A D ^ |   %v < distance from road\n\n", "Y/X", xLine)
}

func prepareMatrix(pl *pipeline) map[uint][]string {
	matrix := make(map[uint][]string, 0)
	var maxDist uint = 0
	for _, d := range pl.houses {
		if d > maxDist {
			maxDist = d
		}
	}

	for i = 0; i < uint(len(pl.houses)); i++ {
		arr := make([]string, maxDist+1)
		var j uint
		for j = 0; j <= maxDist; j++ {
			arr[j] = " "
		}
		matrix[i] = arr
	}
	pl.maxDist = maxDist
	return matrix
}
