package pipeline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	houses   []uint
	excepted uint
}

func TestOne(t *testing.T) {
	tcs := []test{
		{
			houses:   []uint{0, 0, 0, 0, 0},
			excepted: 0,
		},
		{
			houses:   []uint{10, 10, 10, 10, 10},
			excepted: 10,
		},
		{
			houses:   []uint{7, 10, 5, 2},
			excepted: 5,
		},
		{
			houses:   []uint{1, 6, 8, 1, 100, 100, 100, 100},
			excepted: 8,
		},
		{
			houses:   []uint{2, 3, 1, 5},
			excepted: 2,
		},
		{
			houses:   []uint{100, 0, 1000, 500, 0},
			excepted: 466,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := calculateLocation(tc.houses)
			require.Equal(t, tc.excepted, res)
		})
	}
}
