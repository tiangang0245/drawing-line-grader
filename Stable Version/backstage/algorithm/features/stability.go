package features

import (
	"math"
	"line-quality-evaluator/types"
)

func CalculateStability(points []types.StrokePoint) float64 {
	if len(points) < 3 {
		return 1.0
	}

	n := float64(len(points))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	for _, p := range points {
		sumX += p.X
		sumY += p.Y
		sumXY += p.X * p.Y
		sumX2 += p.X * p.X
	}

	denominator := n*sumX2 - sumX*sumX
	if denominator == 0 {
		return 1.0
	}

	slope := (n*sumXY - sumX*sumY) / denominator
	intercept := (sumY - slope*sumX) / n

	totalDistance := 0.0
	for _, p := range points {
		expectedY := slope*p.X + intercept
		distance := math.Abs(p.Y - expectedY)
		totalDistance += distance
	}

	avgDistance := totalDistance / n
	stability := math.Max(0, 1-avgDistance*10)
	return math.Min(1, stability)
}
