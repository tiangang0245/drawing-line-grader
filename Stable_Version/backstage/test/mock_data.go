package test

import (
	"math"
	"math/rand"
	"line-quality-evaluator/types"
)

func CreateMockLine(startX, startY, endX, endY float64, numPoints int) []types.StrokePoint {
	points := make([]types.StrokePoint, numPoints+1)
	startTime := int64(0)

	for i := 0; i <= numPoints; i++ {
		t := float64(i) / float64(numPoints)
		x := startX + (endX-startX)*t
		y := startY + (endY-startY)*t
		time := startTime + int64(i*10)

		points[i] = types.StrokePoint{
			X: x,
			Y: y,
			T: time,
		}
	}

	return points
}

func CreateMockWavyLine(startX, startY, endX, endY float64, numPoints int, amplitude float64) []types.StrokePoint {
	points := make([]types.StrokePoint, numPoints+1)
	startTime := int64(0)

	for i := 0; i <= numPoints; i++ {
		t := float64(i) / float64(numPoints)
		wave := math.Sin(t*math.Pi*6) * amplitude
		x := startX + (endX-startX)*t
		y := startY + (endY-startY)*t + wave
		time := startTime + int64(i*10)

		points[i] = types.StrokePoint{
			X: x,
			Y: y,
			T: time,
		}
	}

	return points
}

func CreateMockJitteryLine(startX, startY, endX, endY float64, numPoints int, jitterAmount float64) []types.StrokePoint {
	points := make([]types.StrokePoint, numPoints+1)
	startTime := int64(0)

	for i := 0; i <= numPoints; i++ {
		t := float64(i) / float64(numPoints)
		jitterX := (rand.Float64() - 0.5) * jitterAmount * 2
		jitterY := (rand.Float64() - 0.5) * jitterAmount * 2
		x := startX + (endX-startX)*t + jitterX
		y := startY + (endY-startY)*t + jitterY
		time := startTime + int64(i*10)

		points[i] = types.StrokePoint{
			X: x,
			Y: y,
			T: time,
		}
	}

	return points
}
