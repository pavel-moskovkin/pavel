package pipeline

//
// https://gist.github.com/rusdevops/d85340e26aeac720c338874492adf637#file-21645-md
//

// The idea is to gradually move the pipe from the road and calculate overall length from the pipeline
// to the houses. When overall length becomes bigger than on previous iteration, return.
func calculateLocation(houses []uint) (dist uint) {
	// total length from houses to the pipe on prev and current iteration
	var sumBeforeMove, sumAfterMove uint = 0, 0
	// 'isCrossed[i]=false' if the pipe is located between a house[i] and the road
	// 'isCrossed[i]=true' if a house[i] is between the pipe and the road
	isCrossed := make(map[int]bool, 0)

	// init values
	for i, h := range houses {
		sumBeforeMove += h
		isCrossed[i] = false
	}

	// func 'move' used to increment distance from the pipe parallel to the road and the road,
	// and calculate required variables.
	move := func() {
		sumAfterMove = 0
		for i := range houses {
			if !isCrossed[i] {
				// decreasing value in the array is equal to decreasing distance from aa house to the pipe
				houses[i] = houses[i] - 1
				sumAfterMove += houses[i]
				if houses[i] == 0 {
					isCrossed[i] = true
				}
				continue
			}
			// if the house is already between the pipe and the road, instead increment value (distance)
			houses[i] = houses[i] + 1
			sumAfterMove += houses[i]
		}
	}

	var i uint
	for i = 0; true; i++ {
		move()
		// When the assertion is correct, no need to continue.
		// It is often happen that (sumBeforeMove == sumAfterMove), so both distances valid.
		// In that case we can take the closest distance.
		if sumBeforeMove <= sumAfterMove {
			break
		}
		sumBeforeMove = sumAfterMove
	}
	return i
}
