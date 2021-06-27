package pipeline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleLine(t *testing.T) {
	houses := []uint{2, 2, 2, 2}
	pipeline, l := calculateLocation(houses)

	expLen := len(houses) - 1
	require.Len(t, pipeline.pipes, expLen)
	require.Equal(t, expLen, l)
	var i uint
	for i = 0; i < uint(len(pipeline.pipes)); i++ {
		pipe := pipeline.pipes[i]
		require.Equal(t, houses[i], pipe.start.x)
		require.Equal(t, houses[i+1], pipe.finish.x)
		require.Equal(t, i, pipe.start.y)
		require.Equal(t, i+1, pipe.finish.y)
	}
}

type tcs struct {
	houses []uint
	exp    []pipe
}

func TestCoords(t *testing.T) {
	tests := []tcs{
		{
			houses: []uint{1, 2, 3},
			exp: []pipe{
				{start: coords{x: 1, y: 0}, finish: coords{x: 1, y: 1}},
				{start: coords{x: 1, y: 1}, finish: coords{x: 2, y: 1}},
				{start: coords{x: 2, y: 1}, finish: coords{x: 2, y: 2}},
				{start: coords{x: 2, y: 2}, finish: coords{x: 3, y: 2}},
			},
		},
		{
			houses: []uint{5, 3, 0},
			exp: []pipe{
				{start: coords{x: 5, y: 0}, finish: coords{x: 3, y: 0}},
				{start: coords{x: 3, y: 0}, finish: coords{x: 3, y: 1}},
				{start: coords{x: 3, y: 1}, finish: coords{x: 0, y: 1}},
				{start: coords{x: 0, y: 1}, finish: coords{x: 0, y: 2}},
			},
		},
		{
			houses: []uint{1, 4, 2},
			exp: []pipe{
				{start: coords{x: 1, y: 0}, finish: coords{x: 1, y: 1}},
				{start: coords{x: 1, y: 1}, finish: coords{x: 4, y: 1}},
				{start: coords{x: 2, y: 1}, finish: coords{x: 2, y: 2}},
			},
		},
		{
			houses: []uint{2, 4, 1},
			exp: []pipe{
				{start: coords{x: 2, y: 0}, finish: coords{x: 2, y: 1}},
				{start: coords{x: 2, y: 1}, finish: coords{x: 4, y: 1}},
				{start: coords{x: 4, y: 1}, finish: coords{x: 1, y: 1}},
				{start: coords{x: 1, y: 1}, finish: coords{x: 1, y: 2}},
			},
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			location, _ := calculateLocation(tc.houses)
			location.printPipeline()
			for j, p := range location.pipes {
				require.Equal(t, tc.exp[j], p)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	houses := []uint{2, 2, 2, 2}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestPrintSnakeLeftToRight(t *testing.T) {
	houses := []uint{1, 2, 3, 4, 5, 6}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestPrintSnakeRightToLeft(t *testing.T) {
	houses := []uint{6, 5, 4, 3, 2, 1}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestPrintComplicated(t *testing.T) {
	houses := []uint{4, 3, 2, 1, 0, 2, 3, 4, 5, 6}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestPrintComplicated2(t *testing.T) {
	houses := []uint{4, 0, 2, 4, 0, 1, 3, 6}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestPrintComplicated3(t *testing.T) {
	houses := []uint{0, 1, 2, 1, 1, 2, 6, 0, 1}
	pipeline, _ := calculateLocation(houses)
	pipeline.printPipeline()
}

func TestVerifyLocation(t *testing.T) {
	houses := []uint{0, 1, 4, 1}
	pipeline, l := calculateLocation(houses)

	var (
		expLen, expCount = 7, 5
	)
	pipeline.printPipeline()
	require.Equal(t, expLen, l)
	require.Len(t, pipeline.pipes, expCount)
}

func TestVerifyLocation2(t *testing.T) {
	houses := []uint{5, 0, 3, 3, 0}
	pipeline, l := calculateLocation(houses)

	var (
		expLen, expCount = 15, 7
	)
	pipeline.printPipeline()
	require.Equal(t, expLen, l)
	require.Len(t, pipeline.pipes, expCount)
}
