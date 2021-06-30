package pipeline

//
// https://gist.github.com/rusdevops/d85340e26aeac720c338874492adf637#file-21645-md
//

func calculateLocation(houses []uint) (dist uint) {
	var sumBeforeMove, sumAfterMove uint = 0, 0

	isCrossed := make(map[int]bool, 0)
	for i, h := range houses {
		sumBeforeMove += h
		isCrossed[i] = false
	}

	move := func() {
		sumAfterMove = 0
		for i := range houses {
			if !isCrossed[i] {
				houses[i] = houses[i] - 1
				sumAfterMove += houses[i]
				if houses[i] == 0 {
					isCrossed[i] = true
				}
				continue
			}
			houses[i] = houses[i] + 1
			sumAfterMove += houses[i]
		}
	}

	var i uint
	for i = 0; true; i++ {
		move()
		if sumBeforeMove <= sumAfterMove {
			break
		}
		sumBeforeMove = sumAfterMove
	}
	return i
}
